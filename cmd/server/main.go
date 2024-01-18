package main

import (
	"context"
	"errors"
	"fmt"
	refresh_token "github.com/LLM-Tests-Checker/Common-Backend/internal/api/auth/refresh-token"
	sign_in "github.com/LLM-Tests-Checker/Common-Backend/internal/api/auth/sign-in"
	sign_up "github.com/LLM-Tests-Checker/Common-Backend/internal/api/auth/sign-up"
	get_results "github.com/LLM-Tests-Checker/Common-Backend/internal/api/llm/get-results"
	get_statuses "github.com/LLM-Tests-Checker/Common-Backend/internal/api/llm/get-statuses"
	launch_check "github.com/LLM-Tests-Checker/Common-Backend/internal/api/llm/launch-check"
	create_test "github.com/LLM-Tests-Checker/Common-Backend/internal/api/tests/create-test"
	delete_test "github.com/LLM-Tests-Checker/Common-Backend/internal/api/tests/delete-test"
	get_my_tests "github.com/LLM-Tests-Checker/Common-Backend/internal/api/tests/get-my-tests"
	get_test "github.com/LLM-Tests-Checker/Common-Backend/internal/api/tests/get-test"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/tests/mappers"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/jwt"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	config2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/config"
	logger2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/auth"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	llm2 "github.com/LLM-Tests-Checker/Common-Backend/internal/storage/llm"
	tests2 "github.com/LLM-Tests-Checker/Common-Backend/internal/storage/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/storage/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Errorf("godotenv.Load: %s", err)
		os.Exit(1)
	}

	config := config2.ProvideServerConfig()
	ctx := context.Background()
	logger := configureLogger(ctx, config)

	logger.Info("Server is starting")

	serverPort, err := config.GetServerPort()
	if err != nil {
		logrus.Errorf("config.GetServerPort: %s", err)
		os.Exit(1)
	}

	router := configureRouter(logger, ctx, config)

	server := http.Server{
		Addr:              fmt.Sprintf("localhost:%s", serverPort),
		Handler:           router,
		ReadTimeout:       2 * time.Second,
		ReadHeaderTimeout: 500 * time.Millisecond,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		BaseContext: func(listener net.Listener) context.Context {
			return context.WithValue(ctx, logger2.Logger, logger)
		},
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Infof("Server started on port: %s", serverPort)

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("server.ListerAndServer: %s", err)
			close(done)
		}
	}()

	<-done
	logger.Infof("Server is stopping")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		logger.Errorf("server.Shutdown: %s", err)
		os.Exit(1)
	}

	logger.Infof("Server stopped")
}

func configureRouter(
	logger *logrus.Logger,
	ctx context.Context,
	config config2.Server,
) *chi.Mux {
	router := chi.NewRouter()

	router.Use(logger2.LoggingMiddleware)
	router.Use(logger2.InfrastructureMiddleware)
	router.Use(middleware.Recoverer)

	launchEnvironment, err := config.GetEnvironment()
	if err != nil {
		logger.Errorf("config.GetEnvironment: %s", err)
		os.Exit(1)
	}

	mongoUrl, err := config.GetMongoUrl()
	if err != nil {
		logger.Errorf("config.GetMongoUrl: %s", err)
		os.Exit(1)
	}

	mongodbLogLevel := options2.LogLevelInfo
	if launchEnvironment == config2.EnvironmentLocal {
		mongodbLogLevel = options2.LogLevelDebug
	}

	mongoLogOptions := options2.Logger().SetComponentLevel(options2.LogComponentAll, mongodbLogLevel)
	options := options2.Client().
		ApplyURI(mongoUrl).
		SetTimeout(time.Second).
		SetAppName("common-backend").
		SetConnectTimeout(10 * time.Second).
		SetMaxConnecting(10).
		SetMinPoolSize(5).
		SetRetryReads(true).
		SetMaxConnIdleTime(30 * time.Second).
		SetServerSelectionTimeout(10 * time.Second).
		SetLoggerOptions(mongoLogOptions)

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		logger.Errorf("Can't connect to mongo: %s", err)
		os.Exit(1)
	}

	databaseName, err := config.GetMongoDatabase()
	if err != nil {
		logger.Errorf("config.GetMongoDatabase: %s", err)
		os.Exit(1)
	}

	database := client.Database(databaseName)

	accessTokenLifetime, err := config.GetAccessTokenLifetime()
	refreshTokenLifetime, err := config.GetRefreshTokenLifetime()
	accessTokenSecret, err := config.GetAccessTokenSecret()
	refreshTokenSecret, err := config.GetRefreshTokenSecret()

	jwtConfig := jwt.Config{
		AccessTokenLiveTime:  accessTokenLifetime,
		RefreshTokenLiveTime: refreshTokenLifetime,
		AccessSecretKey:      accessTokenSecret,
		RefreshSecretKey:     refreshTokenSecret,
	}
	jwtComponent := jwt.NewJWTComponent(jwtConfig)

	userStorage := user.NewUserStorage(logger, database)
	testsStorage := tests2.NewTestsStorage(logger, database)
	llmStorage := llm2.NewLLMStorage(logger, database)

	authService := auth.NewAuthService(userStorage, jwtComponent)
	llmService := llm.NewLLMService(testsStorage, llmStorage)
	testsService := tests.NewTestsService(testsStorage)

	userValidator := users.NewValidator()

	testDtoMapper := mappers.NewTestMapper()

	refreshTokenHandler := refresh_token.New(logger, authService)
	signInHandler := sign_in.New(logger, authService, userValidator)
	signUpHandler := sign_up.New(logger, authService, userValidator)

	getLLMResultsHandler := get_results.New(logger, llmService, authService)
	getLLMStatusesHandler := get_statuses.New(logger, llmService, authService)
	launchLLMCheckHandler := launch_check.New(logger, llmService, authService)

	createTestHandler := create_test.New(logger, testsService, testDtoMapper, authService)
	deleteTestHandler := delete_test.New(logger, testsService, authService)
	getMyTestsHandler := get_my_tests.New(logger, testsService, testDtoMapper, authService)
	getTestHandler := get_test.New(logger, testsService, testDtoMapper, authService)

	server := server{
		refreshToken: refreshTokenHandler,
		signIn:       signInHandler,
		signUp:       signUpHandler,
		getResults:   getLLMResultsHandler,
		getStatuses:  getLLMStatusesHandler,
		launchCheck:  launchLLMCheckHandler,
		createTest:   createTestHandler,
		deleteTest:   deleteTestHandler,
		getMyTests:   getMyTestsHandler,
		getTest:      getTestHandler,
	}

	dto.HandlerFromMux(&server, router)

	return router
}

func configureLogger(
	ctx context.Context,
	config config2.Server,
) *logrus.Logger {
	logger := logrus.New()

	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.PrettyPrint = false

	launchEnvironment, err := config.GetEnvironment()
	if err != nil {
		logger.Errorf("config.GetEnvironemnt: %s", err)
		os.Exit(1)
	}

	logger.WithContext(ctx)
	logger.SetReportCaller(true)
	logger.SetFormatter(formatter)
	logger.WithField("environment", launchEnvironment)
	logger.WithField("application", "server")

	return logger
}

type server struct {
	refreshToken *refresh_token.Handler
	signIn       *sign_in.Handler
	signUp       *sign_up.Handler

	getResults  *get_results.Handler
	getStatuses *get_statuses.Handler
	launchCheck *launch_check.Handler

	createTest *create_test.Handler
	deleteTest *delete_test.Handler
	getMyTests *get_my_tests.Handler
	getTest    *get_test.Handler
}

func (s *server) AuthRefreshToken(w http.ResponseWriter, r *http.Request, params dto.AuthRefreshTokenParams) {
	s.refreshToken.AuthRefreshToken(w, r, params)
}

func (s *server) AuthSignIn(w http.ResponseWriter, r *http.Request) {
	s.signIn.AuthSignIn(w, r)
}

func (s *server) AuthSignUp(w http.ResponseWriter, r *http.Request) {
	s.signUp.AuthSignUp(w, r)
}

func (s *server) TestCreate(w http.ResponseWriter, r *http.Request) {
	s.createTest.TestCreate(w, r)
}

func (s *server) TestDelete(w http.ResponseWriter, r *http.Request, testId dto.TestId) {
	s.deleteTest.TestDelete(w, r, testId)
}

func (s *server) TestById(w http.ResponseWriter, r *http.Request, testId dto.TestId) {
	s.getTest.TestById(w, r, testId)
}

func (s *server) LlmLaunch(w http.ResponseWriter, r *http.Request, testId dto.TestId) {
	s.launchCheck.LlmLaunch(w, r, testId)
}

func (s *server) LlmResult(w http.ResponseWriter, r *http.Request, testId dto.TestId) {
	s.getResults.LlmResult(w, r, testId)
}

func (s *server) LlmStatus(w http.ResponseWriter, r *http.Request, testId dto.TestId) {
	s.getStatuses.LlmStatus(w, r, testId)
}

func (s *server) TestsMy(w http.ResponseWriter, r *http.Request, params dto.TestsMyParams) {
	s.getMyTests.TestsMy(w, r, params)
}

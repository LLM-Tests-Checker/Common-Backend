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
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	logger2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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

	ctx := context.Background()
	logger := configureLogger(ctx)

	serverPort, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		serverPort = "8080"
	}

	router := configureRouter(logger)

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

func configureRouter(logger *logrus.Logger) *chi.Mux {
	router := chi.NewRouter()

	router.Use(logger2.RequestTraceIdMiddleware)
	router.Use(logger2.LoggingMiddleware)
	router.Use(middleware.Recoverer)

	refreshTokenHandler := refresh_token.New(logger)
	signInHandler := sign_in.New(logger)
	signUpHandler := sign_up.New(logger)

	getLLMResultsHandler := get_results.New(logger)
	getLLMStatusesHandler := get_statuses.New(logger)
	launchLLMCheckHandler := launch_check.New(logger)

	createTestHandler := create_test.New(logger)
	deleteTestHandler := delete_test.New(logger)
	getMyTestsHandler := get_my_tests.New(logger)
	getTestHandler := get_test.New(logger)

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

func configureLogger(ctx context.Context) *logrus.Logger {
	logger := logrus.New()

	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.PrettyPrint = false

	launchEnvironment, exists := os.LookupEnv("ENVIRONMENT")
	if !exists {
		launchEnvironment = "local"
	}

	logger.WithContext(ctx)
	logger.SetReportCaller(true)
	logger.SetFormatter(formatter)
	logger.WithField("environment", launchEnvironment)

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

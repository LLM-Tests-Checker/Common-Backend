package main

import (
	"context"
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
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
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

	logger.Infof("Server started on port: %s", serverPort)
}

func configureRouter(logger *logrus.Logger) *chi.Mux {
	router := chi.NewRouter()

	swagger, err := dto.GetSwagger()
	if err != nil {
		logrus.Errorf("dto.GetSwagger: %s", err)
	}

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

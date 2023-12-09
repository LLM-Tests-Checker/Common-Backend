package main

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/auth"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logging"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	router := configureRoutes()
	err = http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func configureRoutes() *mux.Router {
	router := mux.NewRouter()

	addAuthRouts(router)
	addTestsRouts(router)
	addLLMRouts(router)

	router.Use(logging.RequestTracingIdInflatingMiddleware)

	return router
}

func addAuthRouts(router *mux.Router) {
	router.
		Methods(http.MethodPost).
		Path(constants.SignUpPath).
		HandlerFunc(auth.SignUpHandler)
	router.
		Methods(http.MethodPost).
		Path(constants.SignInPath).
		HandlerFunc(auth.SignInHandler)
	router.
		Methods(http.MethodPost).
		Path(constants.RefreshAccessTokenPath).
		HandlerFunc(auth.RefreshAccessTokenHandler)
}

func addTestsRouts(router *mux.Router) {
	getMyTestsRouter := router.
		Methods(http.MethodGet).
		Subrouter()
	getMyTestsRouter.
		Path(constants.GetMyTestsPath).
		HandlerFunc(tests.GetMyTestsHandler)
	getMyTestsRouter.Use(auth.AccessTokenValidationMiddleware)

	getTestByIdRouter := router.
		Methods(http.MethodGet).
		Subrouter()
	getTestByIdRouter.
		Path(constants.GetTestByIdPath).
		HandlerFunc(tests.GetTestByIdHandler)
	getTestByIdRouter.Use(auth.AccessTokenValidationMiddleware)

	createTestRouter := router.
		Methods(http.MethodPut).
		Subrouter()
	createTestRouter.
		Path(constants.CreateTestPath).
		HandlerFunc(tests.CreateTestHandler)
	createTestRouter.Use(auth.AccessTokenValidationMiddleware)

	deleteTestRouter := router.
		Methods(http.MethodDelete).
		Subrouter()
	deleteTestRouter.
		Path(constants.DeleteTestByIdPath).
		HandlerFunc(tests.DeleteTestHandler)
	deleteTestRouter.Use(auth.AccessTokenValidationMiddleware)
}

func addLLMRouts(router *mux.Router) {
	launchLLMRouter := router.
		Methods(http.MethodPost).
		Subrouter()
	launchLLMRouter.
		Path(constants.LaunchLLMCheckPath).
		HandlerFunc(llm.LaunchLLMCheckHandler)
	launchLLMRouter.Use(auth.AccessTokenValidationMiddleware)

	llmStatusRouter := router.
		Methods(http.MethodGet).
		Subrouter()
	llmStatusRouter.
		Path(constants.GetLLMCheckStatusPath).
		HandlerFunc(llm.GetLLMCheckStatusHandler)
	llmStatusRouter.Use(auth.AccessTokenValidationMiddleware)

	llmResultRouter := router.
		Methods(http.MethodGet).
		Subrouter()
	llmResultRouter.
		Path(constants.GetLLMCheckResultPath).
		HandlerFunc(llm.GetLLMCheckResultHandler)
	llmResultRouter.Use(auth.AccessTokenValidationMiddleware)
}

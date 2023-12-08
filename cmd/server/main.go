package main

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/auth"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/tests"
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
	router.
		Methods(http.MethodGet).
		Path(constants.GetMyTestsPath).
		HandlerFunc(tests.GetMyTestsHandler)
	router.
		Methods(http.MethodGet).
		Path(constants.GetTestByIdPath).
		HandlerFunc(tests.GetTestByIdHandler)
	router.
		Methods(http.MethodPut).
		Path(constants.CreateTestPath).
		HandlerFunc(tests.CreateTestHandler)
	router.
		Methods(http.MethodDelete).
		Path(constants.DeleteTestByIdPath).
		HandlerFunc(tests.DeleteTestHandler)
}

func addLLMRouts(router *mux.Router) {
	router.
		Methods(http.MethodPost).
		Path(constants.LaunchLLMCheckPath).
		HandlerFunc(llm.LaunchLLMCheckHandler)
	router.
		Methods(http.MethodGet).
		Path(constants.GetLLMCheckStatusPath).
		HandlerFunc(llm.GetLLMCheckStatusHandler)
	router.
		Methods(http.MethodGet).
		Path(constants.GetLLMCheckResultPath).
		HandlerFunc(llm.GetLLMCheckResultHandler)
}

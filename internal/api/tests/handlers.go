package tests

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services"
	"net/http"
)

var (
	testsService *services.TestsService
)

func init() {
	selector := tests.NewSelector()
	updater := tests.NewUpdater()

	testsService = services.NewTestsService(selector, updater)
}

func GetMyTestsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	testsService.GetMyTests()
}

func GetTestByIdHandler(responseWriter http.ResponseWriter, request *http.Request) {
	testsService.GetTestById()
}

func CreateTestHandler(responseWriter http.ResponseWriter, request *http.Request) {
	testsService.CreateTest()
}

func DeleteTestHandler(responseWriter http.ResponseWriter, request *http.Request) {
	testsService.DeleteTestById()
}

package tests

import (
	"encoding/json"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/common"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/tests"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services"
	"github.com/gorilla/mux"
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
	err, currentUserId := http2.GetCurrentUserId(request)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusUnauthorized)
		return
	}

	err, myTests := testsService.GetMyTests(currentUserId)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusBadRequest)
		return
	}

	responseBody := GetMyTestsResponse{
		Tests: make([]TestDTO, len(myTests)),
	}
	for i := range myTests {
		responseBody.Tests[i] = convertTestModelToDto(myTests[i])
	}

	err = json.NewEncoder(responseWriter).Encode(responseBody)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}

func GetTestByIdHandler(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	targetTestId, ok := vars[constants.TestIdPathParameter]
	if !ok {
		apiError := common.ApiError{
			ErrorCode:    constants.ErrorInvalidRequestParameters,
			ErrorMessage: "Missing test id path parameter",
		}
		http2.ReturnApiError(responseWriter, apiError, http.StatusBadRequest)
	}
	err, currentUserId := http2.GetCurrentUserId(request)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusUnauthorized)
		return
	}

	err, targetTest := testsService.GetTestById(currentUserId, targetTestId)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusBadRequest)
		return
	}

	responseBody := GetTestByIdResponse{
		Test: convertTestModelToDto(*targetTest),
	}
	err = json.NewEncoder(responseWriter).Encode(responseBody)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}

func CreateTestHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var requestBody CreateTestRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusBadRequest)
		return
	}
	err, currentUserId := http2.GetCurrentUserId(request)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusUnauthorized)
		return
	}

	createTestData := convertCreateTestDTOToModel(requestBody)
	err, createdTest := testsService.CreateTest(currentUserId, createTestData)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusBadRequest)
		return
	}

	responseBody := CreateTestResponse{
		Test: convertTestModelToDto(*createdTest),
	}
	err = json.NewEncoder(responseWriter).Encode(responseBody)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}

func DeleteTestHandler(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	targetTestId, ok := vars[constants.TestIdPathParameter]
	if !ok {
		apiError := common.ApiError{
			ErrorCode:    constants.ErrorInvalidRequestParameters,
			ErrorMessage: "Missing test id path parameter",
		}
		http2.ReturnApiError(responseWriter, apiError, http.StatusBadRequest)
	}
	err, currentUserId := http2.GetCurrentUserId(request)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusUnauthorized)
		return
	}

	err = testsService.DeleteTestById(currentUserId, targetTestId)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusBadRequest)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}

func convertTestModelToDto(testModel tests.Test) TestDTO {
	testDTO := TestDTO{
		Identifier:  testModel.Identifier,
		Name:        testModel.Name,
		Description: testModel.Description,
		Questions:   make([]TestQuestionDTO, len(testModel.Questions)),
	}

	for i, question := range testModel.Questions {
		answers := make([]QuestionAnswerDTO, len(question.Answers))
		for j := range question.Answers {
			answers[j] = QuestionAnswerDTO{
				Number:    question.Answers[j].Number,
				Text:      question.Answers[j].Text,
				IsCorrect: question.Answers[j].IsCorrect,
			}
		}

		testDTO.Questions[i] = TestQuestionDTO{
			Number:  question.Number,
			Text:    question.Text,
			Answers: answers,
		}
	}

	return testDTO
}

func convertCreateTestDTOToModel(createTest CreateTestRequest) tests.CreateTestData {
	createTestData := tests.CreateTestData{
		Name:        createTest.Name,
		Description: createTest.Description,
		Questions:   nil,
	}

	for i, question := range createTest.Questions {
		answers := make([]tests.CreateQuestionAnswerData, len(question.Answers))
		for j := range question.Answers {
			answers[j] = tests.CreateQuestionAnswerData{
				Number:    question.Answers[j].Number,
				Text:      question.Answers[j].Text,
				IsCorrect: question.Answers[j].IsCorrect,
			}
		}

		createTestData.Questions[i] = tests.CreateTestQuestionData{
			Number:  question.Number,
			Text:    question.Text,
			Answers: answers,
		}
	}

	return createTestData
}

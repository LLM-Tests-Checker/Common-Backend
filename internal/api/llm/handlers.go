package llm

import (
	"encoding/json"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/common"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/llm"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	llmService *services.LLMCheckerService
)

func init() {
	launcher := llm.NewLauncher()
	statusSelector := llm.NewStatusSelector()
	resultSelector := llm.NewResultSelector()

	llmService = services.NewLLMCheckerService(launcher, statusSelector, resultSelector)
}

func LaunchLLMCheckHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var requestBody LaunchLLMCheckRequest
	vars := mux.Vars(request)
	targetTestId, ok := vars[constants.TestIdPathParameter]
	if !ok {
		apiError := common.ApiError{
			ErrorCode:    constants.ErrorInvalidRequestParameters,
			ErrorMessage: "Missing test id path parameter",
		}
		http2.ReturnApiError(responseWriter, apiError, http.StatusBadRequest)
	}
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

	err, launchResult := llmService.LaunchLLMCheck(currentUserId, targetTestId, requestBody.LLMSlug)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusBadRequest)
		return
	}
	responseBody := LaunchLLMCheckResponse{
		LaunchIdentifier: launchResult.Identifier,
	}

	err = json.NewEncoder(responseWriter).Encode(responseBody)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}
func GetLLMCheckStatusHandler(responseWriter http.ResponseWriter, request *http.Request) {
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

	err, status := llmService.GetLLMCHeckStatus(currentUserId, targetTestId)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusBadRequest)
		return
	}

	responseBody := convertCheckStatusToDTO(status)
	err = json.NewEncoder(responseWriter).Encode(responseBody)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}

func GetLLMCheckResultHandler(responseWriter http.ResponseWriter, request *http.Request) {
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

	err, result := llmService.GetLLMCheckResult(currentUserId, targetTestId)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusBadRequest)
		return
	}

	responseBody := convertCheckResultToDTO(result)
	err = json.NewEncoder(responseWriter).Encode(responseBody)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}

func convertCheckStatusToDTO(status *llm.CheckStatus) GetLLMCheckStatusResponse {
	responseBody := GetLLMCheckStatusResponse{
		Statuses: make([]GetLLMCheckStatusValue, len(status.Statuses)),
	}

	for i := range status.Statuses {
		responseBody.Statuses[i] = GetLLMCheckStatusValue{
			LLMSlug: status.Statuses[i].LLMSlug,
			Status:  status.Statuses[i].Status,
		}
	}

	return responseBody
}

func convertCheckResultToDTO(result *llm.CheckResult) GetLLMCheckResultResponse {
	responseBody := GetLLMCheckResultResponse{
		Results: make([]GetLLMCheckResultValue, len(result.ModelResults)),
	}

	for i := range result.ModelResults {
		resultAnswers := result.ModelResults[i].Results
		answers := make([]GetLLMCheckResultLLMAnswer, len(resultAnswers))
		for j := range resultAnswers {
			answers[j].QuestionNumber = resultAnswers[j].QuestionNumber
			answers[j].SelectedAnswerNumber = resultAnswers[j].SelectedAnswerNumber
		}

		responseBody.Results[i] = GetLLMCheckResultValue{
			LLMSlug: result.ModelResults[i].LLMSlug,
			Answers: answers,
		}
	}

	return responseBody
}

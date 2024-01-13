package get_results

import (
	"encoding/json"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	logger      *logrus.Logger
	selector    resultsSelector
	tokenParser tokenParser
}

func New(
	logger *logrus.Logger,
	selector resultsSelector,
	tokenParser tokenParser,
) *Handler {
	return &Handler{
		logger:      logger,
		selector:    selector,
		tokenParser: tokenParser,
	}
}

func (handler *Handler) LlmResult(response http.ResponseWriter, r *http.Request, testId dto.TestId) {
	userId, err := http2.GetUserIdFromAccessToken(r, handler.tokenParser)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	modelResults, err := handler.selector.GetAllResultsForTest(r.Context(), userId, testId)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	responseResults := mapModelResultToResponse(modelResults)

	err = json.NewEncoder(response).Encode(responseResults)
	if err != nil {
		http2.ReturnErrorWithStatusCode(response, http.StatusBadRequest, err)
		return
	}
}

func mapModelResultToResponse(models []llm.ModelTestResult) dto.GetLLMCheckResultResponse {
	mapAnswersFn := func(modelAnswers []llm.ModelTestAnswer) []dto.GetLLMCheckResultLLMAnswer {
		result := make([]dto.GetLLMCheckResultLLMAnswer, len(modelAnswers))

		for i := range models {
			result[i] = dto.GetLLMCheckResultLLMAnswer{
				QuestionNumber:       int(modelAnswers[i].QuestionNumber),
				SelectedAnswerNumber: int(modelAnswers[i].SelectedAnswerNumber),
			}
		}

		return result
	}

	response := dto.GetLLMCheckResultResponse{
		Results: make([]dto.GetLLMCheckResultValue, len(models)),
	}

	for i := range models {
		response.Results[i] = dto.GetLLMCheckResultValue{
			LlmSlug: models[i].LLMSlug,
			Answers: mapAnswersFn(models[i].LLMAnswers),
		}
	}

	return response
}

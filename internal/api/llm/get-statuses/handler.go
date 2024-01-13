package get_statuses

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
	selector    statusSelector
	tokenParser tokenParser
}

func New(
	logger *logrus.Logger,
	selector statusSelector,
	tokenParser tokenParser,
) *Handler {
	return &Handler{
		logger:      logger,
		selector:    selector,
		tokenParser: tokenParser,
	}
}

func (handler *Handler) LlmStatus(response http.ResponseWriter, r *http.Request, testId dto.TestId) {
	userId, err := http2.GetUserIdFromAccessToken(r, handler.tokenParser)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	statuses, err := handler.selector.GetModelStatusesForTest(r.Context(), userId, testId)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	responseStatuses := mapModelStatusesToResponse(statuses)

	err = json.NewEncoder(response).Encode(responseStatuses)
	if err != nil {
		http2.ReturnErrorWithStatusCode(response, http.StatusBadRequest, err)
		return
	}
}

func mapModelStatusesToResponse(statuses []llm.ModelCheckStatus) dto.GetLLMCheckStatusResponse {
	modelStatusToEnumFn := func(status llm.CheckStatus) dto.GetLLMCheckStatusValueStatus {
		switch status {
		case llm.StatusNotStarted:
			return dto.NOTSTARTED
		case llm.StatusInProgress:
			return dto.INPROGRESS
		case llm.StatusCompleted:
			return dto.COMPLETED
		case llm.StatusError:
			return dto.ERROR
		default:
			return dto.UNDEFINED
		}
	}

	response := dto.GetLLMCheckStatusResponse{
		Statuses: make([]dto.GetLLMCheckStatusValue, len(statuses)),
	}

	for i := range statuses {
		response.Statuses[i] = dto.GetLLMCheckStatusValue{
			LlmSlug: statuses[i].LLMSlug,
			Status:  modelStatusToEnumFn(statuses[i].Status),
		}
	}

	return response
}

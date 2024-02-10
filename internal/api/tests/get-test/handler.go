package get_test

import (
	"encoding/json"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	"net/http"
)

type Handler struct {
	logger      logger.Logger
	selector    testSelector
	mapper      testMapper
	tokenParser tokenParser
}

func New(
	logger logger.Logger,
	selector testSelector,
	mapper testMapper,
	tokenParser tokenParser,
) *Handler {
	return &Handler{
		logger:      logger,
		selector:    selector,
		mapper:      mapper,
		tokenParser: tokenParser,
	}
}

func (handler *Handler) TestById(response http.ResponseWriter, r *http.Request, testId dto.TestId) {
	userId, err := http2.GetUserIdFromAccessToken(r, handler.tokenParser)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	test, err := handler.selector.GetTestById(r.Context(), userId, testId)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	responseTest := handler.mapper.MapModelToDto(test)

	err = json.NewEncoder(response).Encode(responseTest)
	if err != nil {
		http2.ReturnErrorWithStatusCode(response, http.StatusBadRequest, err)
		return
	}
}

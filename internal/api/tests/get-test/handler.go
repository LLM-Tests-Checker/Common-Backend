package get_test

import (
	"encoding/json"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	logger   *logrus.Logger
	selector testSelector
	mapper   testMapper
}

func New(
	logger *logrus.Logger,
	selector testSelector,
	mapper testMapper,
) *Handler {
	return &Handler{
		logger:   logger,
		selector: selector,
		mapper:   mapper,
	}
}

func (handler *Handler) TestById(response http.ResponseWriter, r *http.Request, testId dto.TestId) {
	userId, err := http2.GetUserIdFromAccessToken(r)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	test, err := handler.selector.GetTestById(userId, testId)
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

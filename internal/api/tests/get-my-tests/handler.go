package get_my_tests

import (
	"encoding/json"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	logger         *logrus.Logger
	selector       testsSelector
	inputValidator inputValidator
	mapper         testMapper
}

func New(
	logger *logrus.Logger,
	selector testsSelector,
	inputValidator inputValidator,
	mapper testMapper,
) *Handler {
	return &Handler{
		logger:         logger,
		selector:       selector,
		inputValidator: inputValidator,
		mapper:         mapper,
	}
}

func (handler *Handler) TestsMy(response http.ResponseWriter, r *http.Request, params dto.TestsMyParams) {
	userId, err := http2.GetUserIdFromAccessToken(r)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	pageNumber := int32(params.PageNumber)
	pageSize := int32(params.PageSize)
	err = handler.inputValidator.ValidatePagingParameters(pageNumber, pageSize)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	tests, err := handler.selector.GetTestsByAuthorId(userId, pageNumber, pageSize)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	responseTests := make([]dto.Test, len(tests))
	for i := range tests {
		responseTests[i] = handler.mapper.MapModelToDto(&tests[i])
	}

	err = json.NewEncoder(response).Encode(responseTests)
	if err != nil {
		http2.ReturnErrorWithStatusCode(response, http.StatusBadRequest, err)
		return
	}
}

package get_my_tests

import (
	"encoding/json"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	"net/http"
)

type Handler struct {
	logger      logger.Logger
	selector    testsSelector
	mapper      testMapper
	tokenParser tokenParser
}

func New(
	logger logger.Logger,
	selector testsSelector,
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

func (handler *Handler) TestsMy(response http.ResponseWriter, r *http.Request, params dto.TestsMyParams) {
	userId, err := http2.GetUserIdFromAccessToken(r, handler.tokenParser)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	pageNumber := int32(0)
	pageSize := int32(10)
	if params.PageNumber != nil {
		pageNumber = int32(*params.PageNumber)
	}
	if params.PageSize != nil {
		pageSize = int32(*params.PageSize)
	}

	err = validatePagingParameters(pageNumber, pageSize)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	tests, err := handler.selector.GetTestsByAuthorId(r.Context(), userId, pageNumber, pageSize)
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

func validatePagingParameters(pageNumber, pageSize int32) error {
	const minPageNumberValue = 0
	const maxPageNumberValue = 100

	if pageNumber < minPageNumberValue || pageNumber > maxPageNumberValue {
		return error2.NewBackendError(
			error2.InputValidationError,
			"Invalid page number",
			http.StatusBadRequest,
		)
	}

	const minPageSizeValue = 1
	const maxPageSizeValue = 30

	if pageSize < minPageSizeValue || pageSize > maxPageSizeValue {
		return error2.NewBackendError(
			error2.InputValidationError,
			"Invalid page size",
			http.StatusBadRequest,
		)
	}

	return nil
}

package http

import (
	"encoding/json"
	"errors"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/common"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	"net/http"
)

func ReturnApiError(responseWriter http.ResponseWriter, error common.ApiError, statusCode int32) {
	encoder := json.NewEncoder(responseWriter)
	err := encoder.Encode(error)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	defaultStatusCode := error.GetDefaultStatusCode()
	if defaultStatusCode != nil {
		responseWriter.WriteHeader(int(*defaultStatusCode))
	} else {
		responseWriter.WriteHeader(int(statusCode))
	}
}

func ReturnError(responseWriter http.ResponseWriter, err error, statusCode int32) {
	var apiError common.ApiError
	ok := errors.As(err, &apiError)
	if ok {
		ReturnApiError(responseWriter, apiError, statusCode)
		return
	}

	apiError = common.ApiError{
		ErrorCode:    constants.ErrorUndefined,
		ErrorMessage: err.Error(),
	}
	ReturnApiError(responseWriter, apiError, statusCode)
}

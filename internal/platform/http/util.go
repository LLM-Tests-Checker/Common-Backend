package http

import (
	"encoding/json"
	"errors"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ReturnErrorWithStatusCode(response http.ResponseWriter, statusCode int32, err error) {
	backendErr := error2.BackendError{}

	logrus.Errorf("Request returned error: %s", err)

	if errors.As(err, &backendErr) {
		apiError := dto.ApiError{
			ErrorCode:    int(backendErr.Code),
			ErrorMessage: backendErr.Message,
		}
		setErrorResponse(response, apiError, backendErr.StatusCode)
	} else {
		apiError := dto.ApiError{
			ErrorCode:    int(error2.UnknownError),
			ErrorMessage: err.Error(),
		}
		setErrorResponse(response, apiError, statusCode)
	}
}

func ReturnError(response http.ResponseWriter, err error) {
	backendErr := error2.BackendError{}

	logrus.Errorf("Request returned error: %s", err)

	if errors.As(err, &backendErr) {
		apiError := dto.ApiError{
			ErrorCode:    int(backendErr.Code),
			ErrorMessage: backendErr.Message,
		}
		setErrorResponse(response, apiError, backendErr.StatusCode)
	} else {
		apiError := dto.ApiError{
			ErrorCode:    int(error2.UnknownError),
			ErrorMessage: err.Error(),
		}
		setErrorResponse(response, apiError, http.StatusInternalServerError)
	}
}

func GetUserIdFromAccessToken(r *http.Request, tokenParser tokenParser) (users.UserId, error) {
	accessToken := r.Header.Get(AccessTokenHeaderName)
	if accessToken == "" {
		err := error2.NewBackendError(
			error2.InvalidAccessToken,
			"Access token is missing",
			http.StatusUnauthorized,
		)
		return 0, err
	}

	userId, err := tokenParser.ParseUserId(r.Context(), accessToken)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func setErrorResponse(response http.ResponseWriter, apiError dto.ApiError, statusCode int32) {
	response.WriteHeader(int(statusCode))
	err := json.NewEncoder(response).Encode(apiError)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("json.NewEncoder: %s", err)
		return
	}
}

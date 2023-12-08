package auth

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/common"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"net/http"
)

func AccessTokenValidationMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		accessTokenValue := request.Header.Get(constants.AccessTokenHeaderName)
		if accessTokenValue == "" {
			apiError := common.ApiError{
				ErrorCode:    constants.ErrorInvalidAccessToken,
				ErrorMessage: "Access token is missing",
			}
			http2.ReturnApiError(responseWriter, apiError)
			return
		}

		nextHandler.ServeHTTP(responseWriter, request)
	})
}

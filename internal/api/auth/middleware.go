package auth

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/common"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/auth"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"net/http"
)

var (
	tokensValidator auth.TokensValidator
)

func init() {
	tokensValidator = auth.NewTokensValidator()
}

func AccessTokenValidationMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		accessTokenValue := request.Header.Get(constants.AccessTokenHeaderName)
		if accessTokenValue == "" {
			apiError := common.ApiError{
				ErrorCode:    constants.ErrorInvalidAccessToken,
				ErrorMessage: "Access token is missing",
			}
			http2.ReturnApiError(responseWriter, apiError, http.StatusUnauthorized)
			return
		}

		err := tokensValidator.ValidateAccessToken(accessTokenValue)
		if err != nil {
			apiError := common.ApiError{
				ErrorCode:    constants.ErrorInvalidAccessToken,
				ErrorMessage: err.Error(),
			}
			http2.ReturnApiError(responseWriter, apiError, http.StatusUnauthorized)
			return
		}

		nextHandler.ServeHTTP(responseWriter, request)
	})
}

package auth

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/common"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/auth"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"net/http"
)

const UserId = "UserId"

var (
	tokensValidator     auth.TokensValidator
	tokenUserIdProvider auth.TokenUserIdProvider
)

func init() {
	tokensValidator = auth.NewTokensValidator()
	tokenUserIdProvider = auth.NewTokensUserIdProvider()
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

		err, userId := tokenUserIdProvider.ProvideUserId(accessTokenValue)
		if err != nil {
			apiError := common.ApiError{
				ErrorCode:    constants.ErrorInvalidAccessToken,
				ErrorMessage: err.Error(),
			}
			http2.ReturnApiError(responseWriter, apiError, http.StatusUnauthorized)
		}

		newContext := context.WithValue(request.Context(), UserId, userId)
		newRequest := request.WithContext(newContext)

		nextHandler.ServeHTTP(responseWriter, newRequest)
	})
}

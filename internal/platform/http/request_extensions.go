package http

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/common"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/auth"
	"net/http"
)

var unauthorizedStatusCode = int32(http.StatusUnauthorized)

func GetCurrentUserId(request *http.Request, userIdProvider auth.TokenUserIdProvider) (error, int32) {
	accessTokenHeader := request.Header.Get(constants.AccessTokenHeaderName)
	if accessTokenHeader == "" {
		apiError := common.NewApiError(
			constants.ErrorInvalidAccessToken,
			"Access token is missing",
			&unauthorizedStatusCode,
		)
		return apiError, 0
	}

	err, userId := userIdProvider.ProvideUserId(accessTokenHeader)
	if err != nil {
		return err, 0
	}
	return nil, userId
}

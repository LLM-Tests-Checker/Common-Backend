package http

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/auth"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/common"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	"net/http"
)

var unauthorizedStatusCode = int32(http.StatusUnauthorized)

func GetCurrentUserId(request *http.Request) (error, int32) {
	userId := request.Context().Value(auth.UserId)
	if userId == nil {
		apiError := common.NewApiError(
			constants.ErrorInvalidCredentials,
			"Credentials is missed",
			&unauthorizedStatusCode,
		)
		return apiError, 0
	}
	userIdInt, ok := userId.(int32)
	if !ok {
		apiError := common.NewApiError(
			constants.ErrorInvalidCredentials,
			"Credentials is invalid",
			&unauthorizedStatusCode,
		)
		return apiError, 0
	}

	return nil, userIdInt
}

package http

import (
	"context"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"net/http"
)

func ReturnErrorWithStatusCode(response http.ResponseWriter, statusCode int32, err error) {

}

func ReturnError(response http.ResponseWriter, err error) {

}

type tokenParser interface {
	ParseUserId(ctx context.Context, accessToken string) (users.UserId, error)
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

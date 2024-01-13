package http

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"net/http"
)

func ReturnErrorWithStatusCode(response http.ResponseWriter, statusCode int32, err error) {

}

func ReturnError(response http.ResponseWriter, err error) {

}

type tokenParser interface {
	ParseUserId(accessToken string) (users.UserId, error)
}

func GetUserIdFromAccessToken(r *http.Request) (users.UserId, error) {

}

package delete_test

import (
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	logger      *logrus.Logger
	deleter     testDeleter
	tokenParser tokenParser
}

func New(
	logger *logrus.Logger,
	deleter testDeleter,
	tokenParser tokenParser,
) *Handler {
	return &Handler{
		logger:      logger,
		deleter:     deleter,
		tokenParser: tokenParser,
	}
}

func (handler *Handler) TestDelete(response http.ResponseWriter, r *http.Request, testId dto.TestId) {
	userId, err := http2.GetUserIdFromAccessToken(r, handler.tokenParser)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	err = handler.deleter.DeleteTest(userId, testId)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	response.WriteHeader(http.StatusOK)
}

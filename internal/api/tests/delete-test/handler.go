package delete_test

import (
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	logger *logrus.Logger
}

func New(
	logger *logrus.Logger,
) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (handler *Handler) TestDelete(w http.ResponseWriter, r *http.Request, testId dto.TestId) {

}

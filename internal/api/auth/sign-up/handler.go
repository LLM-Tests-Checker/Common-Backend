package sign_up

import (
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

func (handler *Handler) AuthSignUp(w http.ResponseWriter, r *http.Request) {

}

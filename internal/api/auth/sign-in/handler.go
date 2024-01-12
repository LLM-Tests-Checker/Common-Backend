package sign_in

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

func (handler *Handler) AuthSignIn(w http.ResponseWriter, r *http.Request) {

}

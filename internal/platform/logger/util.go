package logger

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetLogger(request *http.Request) *logrus.Logger {
	logger := request.Context().Value(Logger)
	return logger.(*logrus.Logger)
}

package logging

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func RequestLoggingMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		logrus.Infof("Request: %s %s", request.Method, request.RequestURI)

		responseWriter.Header().Set("Content-Type", "application/json")
		nextHandler.ServeHTTP(responseWriter, request)
	})
}

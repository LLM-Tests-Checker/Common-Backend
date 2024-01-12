package logger

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		logger := GetLogger(r)

		logger.Infof("Request started: %s %s", r.Method, r.RequestURI)

		next.ServeHTTP(w, r)

		logger.Infof("Request finished: %s %s", r.Method, r.RequestURI)
	}

	return http.HandlerFunc(fn)
}

func RequestTraceIdMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestTraceId := uuid.New().String()
		newContext := context.WithValue(r.Context(), RequestTraceID, requestTraceId)

		logger := GetLogger(r)
		logger.WithField(RequestTraceID, requestTraceId)
		newContext = context.WithValue(r.Context(), Logger, logger)

		newRequest := r.WithContext(newContext)

		next.ServeHTTP(w, newRequest)
	}

	return http.HandlerFunc(fn)
}

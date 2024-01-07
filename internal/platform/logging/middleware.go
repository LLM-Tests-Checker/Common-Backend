package logging

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

const RequestTraceId = "RequestTraceId"

func RequestLoggingMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		logrus.Infof("Request: %s %s", request.Method, request.RequestURI)

		responseWriter.Header().Set("Content-Type", "application/json")

		requestId := uuid.New()
		newContext := context.WithValue(request.Context(), RequestTraceId, requestId.String())
		newRequest := request.WithContext(newContext)

		responseWriter.Header().Set(constants.RequestTraceIdHeaderName, requestId.String())

		nextHandler.ServeHTTP(responseWriter, newRequest)
	})
}

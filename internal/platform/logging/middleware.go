package logging

import "net/http"

func RequestTracingIdInflatingMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")
		nextHandler.ServeHTTP(responseWriter, request)
	})
}

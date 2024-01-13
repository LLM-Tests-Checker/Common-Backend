package logger

import (
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

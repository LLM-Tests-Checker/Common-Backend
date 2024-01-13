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

func InfrastructureMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			w.Header().Set("Content-Type", "application/json")
			err := r.Body.Close()
			if err != nil {
				logger := GetLogger(r)
				logger.Errorf("r.Body.Close: %s", err)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

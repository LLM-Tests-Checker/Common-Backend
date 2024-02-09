package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

var (
	requestsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_millis",
			Help: "Duration of HTTP requests in milliseconds",
		},
		[]string{"path"},
	)
)

func CommonMetricsMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			path := fmt.Sprintf("%s_%s", r.URL.Path, r.Method)
			elapsed := time.Since(start).Milliseconds()

			requestsTotal.WithLabelValues(path).Set(1)
			requestDuration.WithLabelValues(path).Observe(float64(elapsed))
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests processed",
		},
		[]string{"method", "endpoint"},
	)

	httpRequestErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_errors_total",
			Help: "Total number of HTTP request errors",
		},
		[]string{"method", "endpoint"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

// InstrumentMiddleware предоставляет промежуточную логику для сбора метрик запросов
func InstrumentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start).Seconds()

		endpoint := r.URL.Path
		method := r.Method

		httpRequestsTotal.WithLabelValues(method, endpoint).Inc()
		httpRequestDuration.WithLabelValues(method, endpoint).Observe(duration)
	})
}

// RecordError увеличивает счетчик ошибок для заданных метода и конечной точки
func RecordError(method, endpoint string) {
	httpRequestErrors.WithLabelValues(method, endpoint).Inc()
}

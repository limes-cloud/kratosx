package middleware

import (
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "server requests duration(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"kind", "operation"})

	metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "the total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})
)

func Metrics(enable bool) middleware.Middleware {
	if !enable {
		return nil
	}
	return metrics.Server(
		metrics.WithSeconds(prom.NewHistogram(metricSeconds)),
		metrics.WithRequests(prom.NewCounter(metricRequests)),
	)
}

package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"go.opentelemetry.io/otel"

	"github.com/limes-cloud/kratosx/library/env"
)

func Metrics(enable bool) middleware.Middleware {
	if !enable {
		return nil
	}

	meter := otel.Meter(env.Instance().AppName())
	requests, err := metrics.DefaultRequestsCounter(meter, metrics.DefaultServerRequestsCounterName)
	if err != nil {
		return nil
	}
	seconds, err := metrics.DefaultSecondsHistogram(meter, metrics.DefaultServerSecondsHistogramName)
	if err != nil {
		return nil
	}

	return func(handler middleware.Handler) middleware.Handler {
		return metrics.Server(
			metrics.WithSeconds(seconds),
			metrics.WithRequests(requests),
		)(handler)
	}
}

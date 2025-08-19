package middleware

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"

	ec "github.com/limes-cloud/kratosx/config"
)

func Tracer(name string, conf *ec.Tracing) middleware.Middleware {
	if conf == nil {
		conf = &ec.Tracing{}
	}
	provider, err := newTracerProvider(name, conf)
	if err != nil {
		panic(err)
	}
	propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	return tracing.Server(
		tracing.WithTracerProvider(provider),
		tracing.WithPropagator(propagator),
	)
}

func newTracerProvider(app string, conf *ec.Tracing) (trace.TracerProvider, error) {
	serviceName := app
	timeout := time.Second * 10

	if conf.Timeout != 0 {
		timeout = conf.Timeout
	}

	var sampler sdktrace.Sampler
	if conf.SampleRatio == nil {
		sampler = sdktrace.AlwaysSample()
	} else {
		sampler = sdktrace.TraceIDRatioBased(float64(*conf.SampleRatio))
	}

	// attributes for all requests
	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
	)

	providerOptions := []sdktrace.TracerProviderOption{
		sdktrace.WithSampler(sampler),
		sdktrace.WithResource(resources),
	}

	var options []otlptracehttp.Option
	if conf.HttpEndpoint != "" {
		options = append(options, otlptracehttp.WithEndpoint(conf.HttpEndpoint))

		if conf.Timeout != 0 {
			options = append(options, otlptracehttp.WithTimeout(timeout))
		}

		if conf.Insecure != nil && !*conf.Insecure {
			options = append(options, otlptracehttp.WithInsecure())
		}

		client := otlptracehttp.NewClient(
			options...,
		)

		exporter, err := otlptrace.New(context.Background(), client)
		if err != nil {
			return nil, err
		}
		providerOptions = append(providerOptions, sdktrace.WithBatcher(exporter))
	}

	return sdktrace.NewTracerProvider(providerOptions...), nil
}

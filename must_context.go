package kratosx

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"go.opentelemetry.io/otel/trace"
)

type ContextOption struct {
	Trace string
	Span  string
}

type ContextOptionFunc func(*ContextOption)

// WithTrace 主动设置trace信息
func WithTrace(trace string, span string) ContextOptionFunc {
	return func(o *ContextOption) {
		o.Trace = trace
		o.Span = span
	}
}

// MustContext returns the Transport value stored in ctx, if any.
func MustContext(c context.Context, opts ...ContextOptionFunc) Context {
	o := &ContextOption{}
	for _, opt := range opts {
		opt(o)
	}

	if o.Trace != "" {
		c = withTraceContext(c, o.Trace, o.Span)
	}

	app, _ := kratos.FromContext(c)
	return &ctx{
		Context: c,
		AppInfo: app,
	}
}

func withTraceContext(ctx context.Context, t string, s string) context.Context {
	tid, err := trace.TraceIDFromHex(t)
	if err != nil {
		return ctx
	}

	sid, err := trace.SpanIDFromHex(s)
	if err != nil {
		return ctx
	}

	// 创建一个新的SpanContext，使用已知的Trace ID和Span ID
	sc := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: tid,
		SpanID:  sid,
		Remote:  true,
	})

	// 创建一个包含新SpanContext的context
	return trace.ContextWithSpanContext(ctx, sc)
}

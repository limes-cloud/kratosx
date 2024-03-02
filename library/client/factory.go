package client

import (
	"context"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/p2c"

	"github.com/limes-cloud/kratosx/config"
)

// Factory is returns service client.
type Factory func(*config.Client) (Client, error)

type Option func(*options)

type options struct {
	pickerBuilder selector.Builder
}

func WithPickerBuilder(in selector.Builder) Option {
	return func(o *options) {
		o.pickerBuilder = in
	}
}

// NewFactory new a client factory.
func NewFactory(r registry.Discovery, opts ...Option) Factory {
	o := &options{
		pickerBuilder: p2c.NewBuilder(),
	}
	for _, opt := range opts {
		opt(o)
	}
	return func(endpoint *config.Client) (Client, error) {
		picker := o.pickerBuilder.Build()
		ctx, cancel := context.WithCancel(context.Background())
		applier := &nodeApplier{
			cancel:   cancel,
			endpoint: endpoint,
			registry: r,
			picker:   picker,
		}
		if err := applier.apply(ctx); err != nil {
			return nil, err
		}
		return newClient(applier, picker), nil
	}
}

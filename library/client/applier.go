package client

import (
	"context"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/selector"

	"github.com/limes-cloud/kratosx/config"
)

type nodeApplier struct {
	// nolint
	canceled int64
	cancel   context.CancelFunc
	endpoint *config.Client
	registry registry.Discovery
	picker   selector.Selector
}

func (na *nodeApplier) apply(_ context.Context) error {
	var nodes []selector.Node
	for _, backend := range na.endpoint.Backends {
		switch na.endpoint.Type {
		case DIRECT:
			nodes = append(nodes, &node{
				address:  backend.Target,
				weight:   backend.Weight,
				metadata: map[string]string{},
			})
			na.picker.Apply(nodes)
		default:
			panic("unknown client type: " + na.endpoint.Type)
		}
	}
	return nil
}

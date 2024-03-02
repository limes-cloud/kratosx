package client

import (
	"context"

	midmetadata "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	ggrpc "google.golang.org/grpc"
)

type client struct {
	applier  *nodeApplier
	selector selector.Selector
}

type Client interface {
	Conn(ctx context.Context) (*ggrpc.ClientConn, error)
}

func newClient(applier *nodeApplier, selector selector.Selector) Client {
	return &client{
		applier:  applier,
		selector: selector,
	}
}

func (c *client) Close() error {
	return nil
}

func (c *client) Conn(ctx context.Context) (*ggrpc.ClientConn, error) {
	if c.applier.endpoint.Type == DIRECT {
		return c.connByDirect(ctx)
	}
	return c.connByDiscovery(ctx)
}

func (c *client) connByDirect(ctx context.Context) (*ggrpc.ClientConn, error) {
	n, done, err := c.selector.Select(ctx)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.DialInsecure(ctx,
		grpc.WithEndpoint(n.Address()),
		grpc.WithMiddleware(midmetadata.Client()),
		grpc.WithTimeout(c.applier.endpoint.Timeout),
	)
	if err != nil {
		done(ctx, selector.DoneInfo{Err: err})
		return nil, err
	}
	return conn, nil
}

func (c *client) connByDiscovery(ctx context.Context) (*ggrpc.ClientConn, error) {
	return grpc.DialInsecure(ctx,
		grpc.WithEndpoint(DISCOVERY+":///"+c.applier.endpoint.Server),
		grpc.WithMiddleware(midmetadata.Client()),
		grpc.WithTimeout(c.applier.endpoint.Timeout),
		grpc.WithDiscovery(c.applier.registry),
	)
}

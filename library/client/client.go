package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"

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

	var callOpts []ggrpc.CallOption
	endpoint := c.applier.endpoint
	if endpoint.MaxRecvSize != 0 {
		callOpts = append(callOpts, ggrpc.MaxCallRecvMsgSize(endpoint.MaxRecvSize*1024*1024))
	}
	if endpoint.MaxSendSize != 0 {
		callOpts = append(callOpts, ggrpc.MaxCallSendMsgSize(endpoint.MaxRecvSize*1024*1024))
	}

	opts := c.options()
	opts = append(opts,
		grpc.WithEndpoint(n.Address()),
		grpc.WithOptions(
			ggrpc.WithDefaultCallOptions(callOpts...),
		),
	)
	conn, err := grpc.DialInsecure(ctx, opts...)
	if err != nil {
		done(ctx, selector.DoneInfo{Err: err})
		return nil, err
	}
	return conn, nil
}

func (c *client) connByDiscovery(ctx context.Context) (*ggrpc.ClientConn, error) {
	opts := c.options()
	opts = append(opts,
		grpc.WithEndpoint(DISCOVERY+":///"+c.applier.endpoint.Server),
		grpc.WithDiscovery(c.applier.registry),
	)
	return grpc.DialInsecure(ctx, opts...)
}

func (c *client) options() []grpc.ClientOption {
	opts := []grpc.ClientOption{
		grpc.WithMiddleware(middlewares(c.applier.endpoint)...),
		grpc.WithTimeout(c.applier.endpoint.Timeout),
	}

	// tls
	if c.applier.endpoint.TLS != nil {
		cp := x509.NewCertPool()
		if cp.AppendCertsFromPEM([]byte(c.applier.endpoint.TLS.Ca)) {
			tlsConf := &tls.Config{ServerName: c.applier.endpoint.TLS.Name, RootCAs: cp}
			opts = append(opts, grpc.WithTLSConfig(tlsConf))
		}
	}
	return opts
}

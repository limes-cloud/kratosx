package client

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	kratosRegistry "github.com/go-kratos/kratos/v2/registry"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/registry"
)

var ins = make(map[string]Client)

const (
	DIRECT    = "direct"
	DISCOVERY = "discovery"
)

// Init 初始化服务
func Init(dsn *string, clients []*config.Client, _ config.Watcher) {
	factory := NewFactory(makeDiscovery(dsn))
	for _, clint := range clients {
		if clint.Timeout == 0 {
			clint.Timeout = time.Second * 10
		}
		c, err := factory(clint)
		if err != nil {
			panic("grpc client init error " + err.Error())
		}
		ins[clint.Server] = c
	}
}

func Get(name string) Client {
	return ins[name]
}

func makeDiscovery(dsn *string) kratosRegistry.Discovery {
	if dsn == nil || *dsn == "" {
		return nil
	}
	d, err := registry.Create(*dsn)
	if err != nil {
		log.Fatalf("failed to create discovery: %v", err)
	}
	return d
}

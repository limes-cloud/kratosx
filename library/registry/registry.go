package registry

import (
	"errors"
	"fmt"
	"net/url"

	kratosRegistry "github.com/go-kratos/kratos/v2/registry"
)

type Interface interface {
	kratosRegistry.Registrar
	kratosRegistry.Discovery
}

type Factory func(dsn *url.URL) (Interface, error)

// Registry is the interface for callers to get registered middleware.
type Registry interface {
	Register(name string, factory Factory)
	Create(registryDSN string) (Interface, error)
}

var globalRegistry = NewRegistry()

type registry struct {
	store map[string]Factory
}

// NewRegistry returns a new middleware registry.
func NewRegistry() Registry {
	return &registry{
		store: map[string]Factory{},
	}
}

func (d *registry) Register(name string, factory Factory) {
	d.store[name] = factory
}

func (d *registry) Create(registryDSN string) (Interface, error) {
	if registryDSN == "" {
		return nil, errors.New("registryDSN is empty")
	}

	dsn, err := url.Parse(registryDSN)
	if err != nil {
		return nil, fmt.Errorf("parse registryDSN error: %s", err)
	}

	factory, ok := d.store[dsn.Scheme]
	if !ok {
		return nil, fmt.Errorf("registry %s has not been registered", dsn.Scheme)
	}

	impl, err := factory(dsn)
	if err != nil {
		return nil, fmt.Errorf("create registry error: %s", err)
	}
	return impl, nil
}

// Register registers one registry.
func Register(name string, factory Factory) {
	globalRegistry.Register(name, factory)
}

// Create instantiates a registry based on `registryDSN`.
func Create(registryDSN string) (Interface, error) {
	return globalRegistry.Create(registryDSN)
}

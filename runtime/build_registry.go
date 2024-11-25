/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package runtime implements the functions, types, and interfaces for the module.
package runtime

import (
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/runtime/registry"
)

type (
	// registryBuildRegistry is an interface that defines a method for registering a RegistryBuilder.
	registryBuildRegistry interface {
		RegisterRegistryBuilder(name string, registryBuilder RegistryBuilder)
	}
	// RegistryBuilder is an interface that defines methods for creating a Discovery and a Registrar.
	RegistryBuilder interface {
		NewRegistrar(cfg *configv1.Registry) (registry.Registrar, error)
		NewDiscovery(cfg *configv1.Registry) (registry.Discovery, error)
	}
)

// RegistrarBuildFunc is a function type that takes a *config.RegistryConfig and returns a registry.Registrar and an error.
type RegistrarBuildFunc func(cfg *configv1.Registry) (registry.Registrar, error)

// NewRegistrar is a method that calls the RegistrarBuildFunc with the given config.
func (fn RegistrarBuildFunc) NewRegistrar(cfg *configv1.Registry) (registry.Registrar, error) {
	return fn(cfg)
}

// DiscoveryBuildFunc is a function type that takes a *config.RegistryConfig and returns a registry.Discovery and an error.
type DiscoveryBuildFunc func(cfg *configv1.Registry) (registry.Discovery, error)

// NewDiscovery is a method that calls the DiscoveryBuildFunc with the given config.
func (fn DiscoveryBuildFunc) NewDiscovery(cfg *configv1.Registry) (registry.Discovery, error) {
	return fn(cfg)
}

// registryWrap is a struct that embeds RegistrarBuildFunc and DiscoveryBuildFunc.
type registryWrap struct {
	RegistrarBuildFunc
	DiscoveryBuildFunc
}

// NewRegistrar creates a new Registrar object based on the given RegistryConfig.
func (b *builder) NewRegistrar(cfg *configv1.Registry) (registry.Registrar, error) {
	b.registryMux.RLock()
	defer b.registryMux.RUnlock()
	registryBuilder, ok := b.registries[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return registryBuilder.NewRegistrar(cfg)
}

// NewDiscovery creates a new Discovery object based on the given RegistryConfig.
func (b *builder) NewDiscovery(cfg *configv1.Registry) (registry.Discovery, error) {
	b.registryMux.RLock()
	defer b.registryMux.RUnlock()
	registryBuilder, ok := b.registries[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return registryBuilder.NewDiscovery(cfg)
}

// RegisterRegistryBuilder registers a new RegistryBuilder with the given name.
func (b *builder) RegisterRegistryBuilder(name string, registryBuilder RegistryBuilder) {
	b.registryMux.Lock()
	defer b.registryMux.Unlock()
	b.registries[name] = registryBuilder
}

// RegisterRegistryFunc registers a new RegistryBuilder with the given name and functions.
func (b *builder) RegisterRegistryFunc(name string, registryBuilder RegistrarBuildFunc, discoveryBuilder DiscoveryBuildFunc) {
	b.RegisterRegistryBuilder(name, &registryWrap{
		RegistrarBuildFunc: registryBuilder,
		DiscoveryBuildFunc: discoveryBuilder,
	})
}

// _ is a blank identifier that is used to satisfy the interface requirement for RegistryBuilder.
var _ RegistryBuilder = &registryWrap{}
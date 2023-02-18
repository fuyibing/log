// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package adapters

import (
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/formatters"
	"sync"
)

type (
	AdapterManager interface {
		// Get
		// registry from manager.
		Get(name string) func() AdapterRegistry

		// Set
		// registry into manager.
		//
		// Override if exists. You must restart client if used
		// registry changed.
		Set(name string, callee func() AdapterRegistry) AdapterManager
	}

	// AdapterRegistry
	// interface for customer adapter.
	AdapterRegistry interface {
		// Logs
		// write lines to target storage / device.
		Logs(lines ...*base.Line) error

		// SetFormatter
		// specify log content formatter handler.
		SetFormatter(formatter formatters.Formatter)
	}

	adapter struct {
		mu         *sync.RWMutex
		registries map[string]func() AdapterRegistry
	}
)

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *adapter) Get(name string) func() AdapterRegistry {
	o.mu.RLock()
	defer o.mu.RUnlock()
	if v, ok := o.registries[name]; ok {
		return v
	}
	return nil
}

func (o *adapter) Set(name string, callee func() AdapterRegistry) AdapterManager {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.registries[name] = callee
	return o
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *adapter) init() *adapter {
	o.mu = &sync.RWMutex{}
	o.registries = make(map[string]func() AdapterRegistry)
	o.initDefaults()
	return o
}

func (o *adapter) initDefaults() {
	for name, call := range adapterContainers {
		o.Set(name, call)
	}
}

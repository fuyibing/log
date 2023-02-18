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
		Get(name string) AdapterRegistry
		Set(name string, registry AdapterRegistry) AdapterManager
	}

	AdapterRegistry interface {
		Logs(lines ...*base.Line) error
		SetFormatter(formatter formatters.Formatter)
	}

	adapter struct {
		mu         *sync.RWMutex
		registries map[string]AdapterRegistry
	}
)

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *adapter) Get(name string) AdapterRegistry {
	o.mu.RLock()
	defer o.mu.RUnlock()
	if v, ok := o.registries[name]; ok {
		return v
	}
	return nil
}

func (o *adapter) Set(name string, registry AdapterRegistry) AdapterManager {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.registries[name] = registry
	return o
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *adapter) init() *adapter {
	o.mu = &sync.RWMutex{}
	o.registries = make(map[string]AdapterRegistry)
	o.initDefaults()
	return o
}

func (o *adapter) initDefaults() {
	for name, call := range adapterDefaults {
		if call != nil {
			o.Set(name, call())
		}
	}
}

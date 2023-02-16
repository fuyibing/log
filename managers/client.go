// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package managers

import (
	"context"
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/conf"
	"sync"
)

type (
	Client struct {
		adapter  adapters.AdapterInterface
		adapters map[conf.Adapter]adapters.AdapterInterface
		mu       *sync.RWMutex
	}
)

func (o *Client) GetAdapterInterface(adapter conf.Adapter) adapters.AdapterInterface {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if v, ok := o.adapters[adapter]; ok {
		return v
	}
	return nil
}

func (o *Client) Init() *Client {
	o.adapters = make(map[conf.Adapter]adapters.AdapterInterface)
	o.mu = &sync.RWMutex{}
	return o
}

func (o *Client) Subscribe() {
	o.mu.Lock()
	defer o.mu.Unlock()

	var (
		i    = 0
		prev adapters.AdapterInterface
	)

	// Reset boot adapter.
	o.adapter = nil

	// Iterate configured adapters.
	for _, adapter := range conf.Config.Adapters {
		if v := adapter.New(); v != nil {
			// Set boot adapter.
			if i++; i == 1 {
				o.adapter = v
			}

			// Set this adapter interface as previous adapter child.
			if prev != nil {
				prev.Child(v)
			}

			// Reset previous adapter interface for next.
			prev = v

			// Update mapper.
			o.adapters[adapter] = v
		}
	}
}

// /////////////////////////////////////////////////////////////
// Exported methods.
// /////////////////////////////////////////////////////////////

func (o *Client) Debug(text string) {
	o.send(nil, conf.Debug, text)
}

func (o *Client) Debugf(text string, args ...interface{}) {
	o.send(nil, conf.Debug, text, args...)
}

func (o *Client) Debugfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Debug, text, args...)
}

func (o *Client) Info(text string) {
	o.send(nil, conf.Info, text)
}

func (o *Client) Infof(text string, args ...interface{}) {
	o.send(nil, conf.Info, text, args...)
}

func (o *Client) Infofc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Info, text, args...)
}

func (o *Client) Warn(text string) {
	o.send(nil, conf.Warn, text)
}

func (o *Client) Warnf(text string, args ...interface{}) {
	o.send(nil, conf.Warn, text, args...)
}

func (o *Client) Warnfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Warn, text, args...)
}

func (o *Client) Error(text string) {
	o.send(nil, conf.Error, text)
}

func (o *Client) Errorf(text string, args ...interface{}) {
	o.send(nil, conf.Error, text, args...)
}

func (o *Client) Errorfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Error, text, args...)
}

func (o *Client) Fatal(text string) {
	o.send(nil, conf.Fatal, text)
}

func (o *Client) Fatalf(text string, args ...interface{}) {
	o.send(nil, conf.Fatal, text, args...)
}

func (o *Client) Fatalfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Fatal, text, args...)
}

// /////////////////////////////////////////////////////////////
// Access methods.
// /////////////////////////////////////////////////////////////

func (o *Client) send(ctx context.Context, level conf.Level, text string, args ...interface{}) {
	if o.adapter != nil {
		o.adapter.Send(nil)
	}
}

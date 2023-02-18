// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package core

import (
	"context"
	"fmt"
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/conf"
	"sync/atomic"
	"time"
)

type (
	Client interface {
		// GetAdapterRegistry
		// return adapter registry.
		GetAdapterRegistry() adapters.AdapterRegistry

		GetBucket() Bucket

		Start()
		Stop()

		ClientUser
	}

	ClientUser interface {
		Debug(text string)
		Debugf(text string, args ...interface{})
		Debugfc(ctx context.Context, text string, args ...interface{})
		Info(text string)
		Infof(text string, args ...interface{})
		Infofc(ctx context.Context, text string, args ...interface{})
		Warn(text string)
		Warnf(text string, args ...interface{})
		Warnfc(ctx context.Context, text string, args ...interface{})
		Error(text string)
		Errorf(text string, args ...interface{})
		Errorfc(ctx context.Context, text string, args ...interface{})
		Fatal(text string)
		Fatalf(text string, args ...interface{})
		Fatalfc(ctx context.Context, text string, args ...interface{})
	}

	client struct {
		ar          adapters.AdapterRegistry
		bucket      Bucket
		concurrency int32
	}
)

func NewClient() Client {
	return (&client{}).
		init()
}

// /////////////////////////////////////////////////////////////
// Exported methods
// /////////////////////////////////////////////////////////////

func (o *client) GetAdapterRegistry() adapters.AdapterRegistry {
	return o.ar
}

func (o *client) GetBucket() Bucket { return o.bucket }

func (o *client) Start() { o.start() }
func (o *client) Stop()  { o.stop() }

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *client) init() *client {
	o.bucket = (&bucket{}).init()
	return o
}

func (o *client) start() {
	// Get adapter from registry.
	if o.ar = adapters.Adapter.Get(conf.Config.GetAdapter()); o.ar == nil {
		panic(fmt.Sprintf("adapter not registered: %s", conf.Config.GetAdapter()))
		return
	}

}

func (o *client) stop() bool {
	// Recall
	// after 3 milliseconds if bucket has remaining lines of
	// running concurrency not completed.
	if concurrency := atomic.LoadInt32(&o.concurrency); concurrency > 0 || o.bucket.Count() > 0 {
		// Open new coroutine
		// to consume bucket.
		if concurrency < conf.Config.GetBatchConcurrency() {
			go o.popn()
		}

		// Wait for a while.
		time.Sleep(time.Millisecond * 3)
		return o.stop()
	}
	return true
}

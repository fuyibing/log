// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package core

import (
	"context"
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/conf"
	"sync/atomic"
	"time"
)

type (
	Client interface {
		// Close
		// log client, block process until all lines completed.
		Close()

		// GetAdapterRegistry
		// return adapter registry.
		GetAdapterRegistry() adapters.AdapterRegistry

		// GetBucket
		// return client lines bucket.
		GetBucket() Bucket

		// PushIntoBucket
		// push log into bucket.
		PushIntoBucket(ctx context.Context, level conf.Level, property map[string]interface{}, text string, args ...interface{})

		// Reset
		// adapter handler.
		Reset()
	}

	client struct {
		ar, ae      adapters.AdapterRegistry
		arc         string
		bucket      Bucket
		cancel      context.CancelFunc
		ctx         context.Context
		concurrency int32
		frequency   Frequency
	}
)

func NewClient() Client {
	return (&client{}).
		init()
}

// /////////////////////////////////////////////////////////////
// Exported methods
// /////////////////////////////////////////////////////////////

func (o *client) Close()                                       { o.close() }
func (o *client) GetAdapterRegistry() adapters.AdapterRegistry { return o.ar }
func (o *client) GetBucket() Bucket                            { return o.bucket }
func (o *client) Reset()                                       { o.reset() }

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *client) close() bool {
	// Cancel
	// frequency context.
	if o.ctx != nil && o.ctx.Err() == nil {
		o.cancel()
	}

	// Recall
	// after 3 milliseconds if bucket has remaining lines of
	// running concurrency not completed.
	if concurrency := atomic.LoadInt32(&o.concurrency); concurrency > 0 || o.bucket.Count() > 0 {
		// Open new coroutine
		// to consume bucket.
		if concurrency < conf.Config.GetBatchConcurrency() {
			go o.PopFromBucket()
		}

		// Wait for a while.
		time.Sleep(time.Millisecond * 3)
		return o.close()
	}
	return true
}

func (o *client) init() *client {
	// Register bucket.
	o.bucket = (&bucket{}).init()

	// Open frequency.
	go func() {
		o.ctx, o.cancel = context.WithCancel(context.Background())
		o.frequency = (&frequency{callback: o.PopFromBucket}).init()
		_ = o.frequency.Processor().Start(o.ctx)
		o.ctx = nil
	}()

	// Register error adapter.
	if call := adapters.Adapter.Get(adapters.AdapterError); call != nil {
		o.ae = call()
	}

	// Set configured adapter.
	o.reset()
	return o
}

func (o *client) reset() {
	if arc := conf.Config.GetAdapter(); arc != o.arc {
		if call := adapters.Adapter.Get(conf.Config.GetAdapter()); call != nil {
			o.ar = call()
			o.arc = arc
		}
	}
}

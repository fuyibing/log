// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package managers

import (
	"context"
	"fmt"
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/adapters/file"
	"github.com/fuyibing/log/v8/adapters/kafka"
	"github.com/fuyibing/log/v8/adapters/term"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/log/v8/formatters"
	"sync/atomic"
	"time"
)

var (
	ErrAdapterUndefined = fmt.Errorf("log adapter undefined")
	ErrNotHealthy       = fmt.Errorf("log client not health")
)

type (
	ClientManager struct {
		adapter    adapters.AdapterManager
		bucket     adapters.Bucket
		cancel     context.CancelFunc
		ch         chan *base.Line
		ct         *time.Ticker
		ctx        context.Context
		formatter  formatters.Formatter
		processor  adapters.Processor
		processing int32

		totalSend, totalPush, totalPop uint64
	}
)

// /////////////////////////////////////////////////////////////
// Exported methods
// /////////////////////////////////////////////////////////////

func (o *ClientManager) Stop()                     { o.stop() }
func (o *ClientManager) Start(ctx context.Context) { o.start(ctx) }

func (o *ClientManager) init() *ClientManager {
	o.formatter = formatters.NewIgnoreFormatter().Format
	o.bucket = adapters.NewBucket()
	o.processor = adapters.NewProcessor("client manager").Callback(
		o.onChannel,
		o.onBucketClean,
		o.onBucketDone,
	).Panic(o.onPanic)

	return o
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *ClientManager) doIgnore(err error, lines ...*base.Line) {
	for _, line := range lines {
		text := o.formatter(line)
		line.Release()
		if err == nil {
			println(text)
		} else {
			println(text, err.Error())
		}
	}
}

func (o *ClientManager) doPop() {
	var (
		concurrency      = atomic.AddInt32(&o.processing, 1)
		count, remaining int
		list             []*base.Line
	)

	// Return
	// if concurrency is greater than configuration.
	if concurrency > conf.Config.BatchConcurrency {
		atomic.AddInt32(&o.processing, -1)
		return
	}

	list, count, remaining = o.bucket.Popn(conf.Config.BatchLimit)
	if count == 0 {
		atomic.AddInt32(&o.processing, -1)
		return
	}

	// Call adapter.
	o.adapter.Log(list...)
	atomic.AddInt32(&o.processing, -1)

	// Continue pop.
	if remaining > 0 {
		o.doPop()
	}
}

func (o *ClientManager) doPush(line *base.Line) {
	if o.bucket.Push(line) >= conf.Config.BatchLimit {
		o.doPop()
	}
}

func (o *ClientManager) send(ctx context.Context, level conf.Level, text string, args ...interface{}) {
	// Acquire line.
	line := base.Pool.AcquireLine()
	line.Ctx = ctx
	line.Level = level
	line.Text = fmt.Sprintf(text, args...)

	// Return if adapter not defined.
	if o.adapter == nil {
		o.doIgnore(ErrAdapterUndefined, line)
		return
	}

	// Return if processor not health.
	if !o.processor.Healthy() {
		o.doIgnore(ErrNotHealthy, line)
		return
	}

	// Catch panic.
	defer func() {
		if v := recover(); v != nil {
			o.doIgnore(fmt.Errorf("%v", v), line)
		}
	}()

	// Send channel signal.
	o.ch <- line
}

func (o *ClientManager) start(ctx context.Context) {
	var (
		curr, last adapters.AdapterManager
	)

	o.adapter = nil

	// Iterate registered.
	for _, adapter := range conf.Config.Adapters {
		// Ignore unknown adapter.
		if curr = o.switchAdapter(adapter); curr == nil {
			continue
		}

		// Register ignore callback on adapter.
		curr.SetIgnore(o.doIgnore)

		// Register current adapter as child of last.
		if last != nil {
			last.SetChild(curr)
		}

		// Update last.
		last = curr

		// Register first adapter.
		if o.adapter == nil {
			o.adapter = curr
		}
	}

	// Start in coroutine.
	if o.adapter != nil {
		// Redirect to background context.
		if ctx == nil {
			ctx = context.Background()
		}

		// Create cancellable context.
		o.ctx, o.cancel = context.WithCancel(ctx)

		// Start processor in coroutine.
		go func(c context.Context) {
			_ = o.processor.Start(c)
		}(o.ctx)

		// Wait for a while.
		time.Sleep(time.Millisecond * 3)
	}

	return

}

func (o *ClientManager) stop() {
	if o.ctx != nil && o.ctx.Err() == nil {
		o.cancel()
	}
	for {
		if o.processor.Stopped() {
			break
		}
		time.Sleep(time.Millisecond * 3)
	}
}

func (o *ClientManager) switchAdapter(name conf.Adapter) adapters.AdapterManager {
	switch name {
	case conf.Term:
		return term.Adapter
	case conf.File:
		return file.Adapter
	case conf.Kafka:
		return kafka.Adapter
	}
	return nil
}

// /////////////////////////////////////////////////////////////
// Processor events
// /////////////////////////////////////////////////////////////

func (o *ClientManager) onChannel(ctx context.Context) (ignored bool) {
	// Build channel.
	o.ch = make(chan *base.Line)
	o.ct = time.NewTicker(time.Duration(conf.Config.BatchDuration) * time.Millisecond)

	// Destroy channel when end.
	defer func() {
		// Unset channel.
		close(o.ch)
		o.ch = nil

		// Clear ticker.
		o.ct.Stop()
		o.ct = nil
	}()

	// Listen channel signal.
	for {
		select {
		case line := <-o.ch:
			go o.doPush(line)
		case <-o.ct.C:
			go o.doPop()
		case <-ctx.Done():
			return
		}
	}
}

func (o *ClientManager) onBucketClean(ctx context.Context) (ignored bool) {
	if o.bucket.Count() > 0 {
		go o.doPop()
		time.Sleep(time.Millisecond * 3)
		return o.onBucketClean(ctx)
	}
	return
}

func (o *ClientManager) onBucketDone(ctx context.Context) (ignored bool) {
	if atomic.LoadInt32(&o.processing) == 0 {
		return
	}

	time.Sleep(time.Millisecond * 3)
	return o.onBucketDone(ctx)
}

func (o *ClientManager) onPanic(_ context.Context, _ interface{}) {}

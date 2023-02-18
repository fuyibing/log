// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package core

import (
	"context"
	"fmt"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
	"sync/atomic"
)

func (o *client) Debug(text string) {
	o.send(nil, conf.Debug, text)
}

func (o *client) Debugf(text string, args ...interface{}) {
	o.send(nil, conf.Debug, text, args...)
}

func (o *client) Debugfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Debug, text, args...)
}

func (o *client) Info(text string) {
	o.send(nil, conf.Info, text)
}

func (o *client) Infof(text string, args ...interface{}) {
	o.send(nil, conf.Info, text, args...)
}

func (o *client) Infofc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Info, text, args...)
}

func (o *client) Warn(text string) {
	o.send(nil, conf.Warn, text)
}

func (o *client) Warnf(text string, args ...interface{}) {
	o.send(nil, conf.Warn, text, args...)
}

func (o *client) Warnfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Warn, text, args...)
}

func (o *client) Error(text string) {
	o.send(nil, conf.Error, text)
}

func (o *client) Errorf(text string, args ...interface{}) {
	o.send(nil, conf.Error, text, args...)
}

func (o *client) Errorfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Error, text, args...)
}

func (o *client) Fatal(text string) {
	o.send(nil, conf.Fatal, text)
}

func (o *client) Fatalf(text string, args ...interface{}) {
	o.send(nil, conf.Fatal, text, args...)
}

func (o *client) Fatalfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Fatal, text, args...)
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *client) popn() {
	// Ignore coroutine if concurrency is greater than
	// configuration.
	if concurrency := atomic.AddInt32(&o.concurrency, 1); concurrency > conf.Config.GetBatchConcurrency() {
		atomic.AddInt32(&o.concurrency, -1)
		return
	}

	// Prepare pop.
	var (
		count  int
		list   []*base.Line
		recall = false
	)

	// End adapter called.
	defer func() {
		atomic.AddInt32(&o.concurrency, -1)

		// Continue call pop
		// until bucket is empty.
		if recall {
			o.popn()
		}
	}()

	// Pop progress.
	if list, _, count, _ = o.bucket.Popn(conf.Config.GetPid()); count > 0 {
		recall = true

		// Send to adapter.
		if err := o.ar.Logs(list...); err == nil {
			_ = o.are.Logs(list...)
		}

		// Release lines.
		o.release(list)
	}
}

func (o *client) release(lines []*base.Line) {
	for _, line := range lines {
		line.Release()
	}
}

func (o *client) send(ctx context.Context, level conf.Level, text string, args ...interface{}) {
	// Acquire
	// line from pool.
	line := base.Pool.AcquireLine().WithContext(ctx)
	line.Level = level
	line.Text = fmt.Sprintf(text, args...)

	// Send to adapter.
	if total, count := o.bucket.Push(line); count > 0 && total >= conf.Config.GetBatchLimit() {
		go o.popn()
	}
}

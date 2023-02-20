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

// /////////////////////////////////////////////////////////////
// Lines operators
// /////////////////////////////////////////////////////////////

func (o *client) CallAdapter(lines ...*base.Line) {
	var (
		err error
	)

	// Call error adapter.
	defer func() {
		// Catch panic.
		if v := recover(); v != nil {
			err = fmt.Errorf("v")
		}

		// Call error handler if send failed.
		if err != nil && o.ae != nil {
			_ = o.ae.Logs(lines...)
		}

		// Release all lines after operated.
		o.ReleaseLines(lines)
	}()

	if o.ar == nil {
		err = fmt.Errorf("unknown adapter registry")
	} else {
		err = o.ar.Logs(lines...)
	}
}

func (o *client) PopFromBucket() {
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
		// Revert concurrency then continue call pop
		// until bucket is empty.
		atomic.AddInt32(&o.concurrency, -1)

		if recall {
			o.PopFromBucket()
		}
	}()

	// Pop progress.
	if list, _, count, _ = o.bucket.Popn(conf.Config.GetBatchLimit()); count > 0 {
		recall = true
		o.CallAdapter(list...)
	}
}

func (o *client) PushIntoBucket(ctx context.Context, level conf.Level, property map[string]interface{}, text string, args ...interface{}) {
	// Acquire
	// line from pool.
	line := base.Pool.AcquireLine().WithContext(ctx)
	line.Property = property
	line.Level = level
	line.Text = fmt.Sprintf(text, args...)
	line.TextParse()

	// SYNC Mode, if ASYNC Disabled.
	if conf.Config.GetAsyncDisabled() {
		o.CallAdapter(line)
		return
	}

	// Send to adapter.
	if total, count := o.bucket.Push(line); count > 0 && total >= conf.Config.GetBatchLimit() {
		go o.PopFromBucket()
	}
}

func (o *client) ReleaseLines(lines []*base.Line) {
	for _, line := range lines {
		line.Release()
	}
}

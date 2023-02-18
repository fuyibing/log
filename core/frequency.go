// author: wsfuyibing <websearch@163.com>
// date: 2023-02-19

package core

import (
	"context"
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/util/v8/process"
	"time"
)

type (
	Frequency interface {
		Processor() process.Processor
	}

	frequency struct {
		callback func()
		process  process.Processor
		ticker   *time.Ticker
	}
)

func (o *frequency) Processor() process.Processor { return o.process }

// /////////////////////////////////////////////////////////////
// Process events
// /////////////////////////////////////////////////////////////

func (o *frequency) OnChannel(ctx context.Context) (ignored bool) {
	for {
		select {
		case <-o.ticker.C:
			go o.callback()
		case <-ctx.Done():
			return
		}
	}
}

func (o *frequency) OnChannelAfter(_ context.Context) (ignored bool) {
	o.ticker.Stop()
	o.ticker = nil
	return
}

func (o *frequency) OnChannelBefore(_ context.Context) (ignored bool) {
	o.ticker = time.NewTicker(time.Duration(conf.Config.GetBatchFrequency()) * time.Millisecond)
	return
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *frequency) init() *frequency {
	o.process = process.New("frequency").
		Callback(o.OnChannelBefore, o.OnChannel, o.OnChannelAfter)
	return o
}

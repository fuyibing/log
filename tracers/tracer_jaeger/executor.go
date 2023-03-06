// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// author: wsfuyibing <websearch@163.com>
// date: 2023-03-05

// Package tracer_jaeger
// 上报到Jaeger.
package tracer_jaeger

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/fuyibing/log/v5/common"
	"github.com/fuyibing/log/v5/configurer"
	"github.com/fuyibing/log/v5/tracers"
	"github.com/fuyibing/util/v8/process"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync/atomic"
	"time"
)

type executor struct {
	bucket     common.Bucket
	formatter  tracers.Formatter
	name       string
	processor  process.Processor
	processing int32
}

func New() tracers.Executor { return (&executor{}).init() }

// /////////////////////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////////////////////

func (o *executor) Processor() process.Processor        { return o.processor }
func (o *executor) Publish(spans ...tracers.Span) error { return o.publish(spans...) }
func (o *executor) SetFormatter(v tracers.Formatter)    { o.formatter = v }

// /////////////////////////////////////////////////////////////////////////////
// Event methods
// /////////////////////////////////////////////////////////////////////////////

func (o *executor) onAfter(ctx context.Context) (ignored bool) {
	cc := atomic.LoadInt32(&o.processing)

	// 处理完成.
	// - 并行降低
	// - 空数据桶.
	if cc == 0 && o.bucket.IsEmpty() {
		return
	}

	// 加大并行.
	if cc < configurer.Config.GetBucketConcurrency() {
		go o.pop()
	}

	// 定时延后.
	time.Sleep(time.Millisecond * 100)
	return o.onAfter(ctx)
}

func (o *executor) onCall(ctx context.Context) (ignored bool) {
	common.InternalInfo("<%s> signal listening", o.name)

	// 定时收取.
	ti := time.NewTicker(time.Duration(configurer.Config.GetBucketFrequency()) * time.Millisecond)

	// 监听信号.
	for {
		select {
		case <-ti.C:
			go o.pop()
		case <-ctx.Done():
			return
		}
	}
}

func (o *executor) onPanic(_ context.Context, v interface{}) {
	common.InternalFatal("<%s> fatal: %v", o.name, v)
}

// /////////////////////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////////////////////

func (o *executor) init() *executor {
	o.bucket = common.NewBucket(configurer.Config.GetBucketCapacity())
	o.formatter = (&formatter{}).init()
	o.name = "tracer.jaeger"
	o.processor = process.New(o.name).
		After(o.onAfter).
		Callback(o.onCall).
		Panic(o.onPanic)

	return o
}

func (o *executor) pop() {
	// 限流控制.
	if cc := atomic.AddInt32(&o.processing, 1); cc > configurer.Config.GetBucketConcurrency() {
		atomic.AddInt32(&o.processing, -1)
		return
	}

	// 取出数据.
	var (
		list []tracers.Span
		redo = false
	)
	if items, _, count := o.bucket.Popn(configurer.Config.GetBucketBatch()); count > 0 {
		list = make([]tracers.Span, 0)
		redo = true
		// 遍历数据.
		for _, item := range items {
			if v, ok := item.(tracers.Span); ok {
				list = append(list, v)
			}
		}
		// 处理跨度.
		if len(list) > 0 {
			if err := o.send(list...); err != nil {
				common.InternalInfo("<%s> send: %v", o.name, err)
			}
		}
	}

	// 恢复并行.
	atomic.AddInt32(&o.processing, -1)
	if redo {
		o.pop()
	}
}

func (o *executor) publish(spans ...tracers.Span) (err error) {
	var total int

	// 健康进程.
	if o.processor.Healthy() {
		// 数据入桶.
		for _, log := range spans {
			if total, err = o.bucket.Add(log); err != nil {
				return
			}
		}

		// 立即消费.
		if total >= configurer.Config.GetBucketBatch() {
			go o.pop()
		}
		return
	}

	// 立即写入.
	return o.send(spans...)
}

func (o *executor) send(spans ...tracers.Span) (err error) {
	if len(spans) == 0 {
		return
	}

	var body []byte

	if body, err = o.formatter.Byte(spans...); err != nil {
		return
	}

	var (
		buf = bytes.NewBuffer(body)
		req = fasthttp.AcquireRequest()
		res = fasthttp.AcquireResponse()
	)

	req.SetRequestURI(configurer.Config.GetJaegerTracer().GetEndpoint())
	req.SetBodyStream(buf, buf.Len())
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType(configurer.Config.GetJaegerTracer().GetContentType())

	// Bind authorization.
	if usr := configurer.Config.GetJaegerTracer().GetUsername(); usr != "" {
		pwd := configurer.Config.GetJaegerTracer().GetPassword()
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(usr+":"+pwd))))
	}

	// Send request.
	err = fasthttp.Do(req, res)
	return
}

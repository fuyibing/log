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
// date: 2023-02-24

// Package tracer_jaeger
// 链路[异步]上传到 Jaeger.
package tracer_jaeger

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/util/v8/process"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/valyala/fasthttp"
)

type Exporter struct {
	formatter  Formatter
	name       string
	processing int32
	processor  process.Processor
}

func New() base.TracerExporter { return (&Exporter{}).init() }

// Processor 获取类进程.
func (o *Exporter) Processor() process.Processor { return o.processor }

// Send 发送日志.
func (o *Exporter) Send(span base.Span) (err error) {
	atomic.AddInt32(&o.processing, 1)
	defer atomic.AddInt32(&o.processing, -1)

	var buf *bytes.Buffer

	if buf, err = o.formatter.Thrift(span); err == nil {
		_ = o.Uploader(buf)
	}
	return
}

// SetFormatter 设置格式化.
func (o *Exporter) SetFormatter(formatter base.TracerFormatter) {}

// Uploader
// 上报日志.
func (o *Exporter) Uploader(buf *bytes.Buffer) (err error) {
	var (
		req = fasthttp.AcquireRequest()
		res = fasthttp.AcquireResponse()
	)

	req.SetRequestURI(conf.Config.GetJaegerTracer().GetEndpoint())
	req.SetBodyStream(buf, buf.Len())
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType(conf.Config.GetJaegerTracer().GetContentType())

	// Bind authorization.
	if usr := conf.Config.GetJaegerTracer().GetUsername(); usr != "" {
		pwd := conf.Config.GetJaegerTracer().GetPassword()
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(usr+":"+pwd))))
	}

	// Send request.
	err = fasthttp.Do(req, res)
	return
}

// /////////////////////////////////////////////////////////////////////////////
// Exporter: events
// /////////////////////////////////////////////////////////////////////////////

func (o *Exporter) init() *Exporter {
	o.formatter = (&formatter{}).init()

	o.name = "exporter.tracer.jaeger"
	o.processor = process.New(o.name).
		After(o.onAfter).
		Before(o.onBefore).
		Callback(o.onCall).
		Panic(o.onPanic)
	return o
}

func (o *Exporter) onAfter(ctx context.Context) (ignored bool) {
	if atomic.LoadInt32(&o.processing) == 0 {
		return
	}

	time.Sleep(time.Millisecond)
	return o.onAfter(ctx)
}

func (o *Exporter) onBefore(_ context.Context) (ignored bool) {
	return
}

func (o *Exporter) onCall(ctx context.Context) (ignored bool) {
	base.InternalInfo("<%s> tracer listening", o.name)

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (o *Exporter) onPanic(_ context.Context, v interface{}) {
	base.InternalError("<%s> %v", o.name, v)
}

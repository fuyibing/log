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

package tracer_jaeger

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/log/v5/cores"
	"github.com/fuyibing/util/v8/process"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/valyala/fasthttp"
)

type (
	// Exporter
	// an exporter component for tracer, publish to jaeger.
	Exporter struct {
		formatter  Formatter
		processing int32
		processor  process.Processor
	}
)

// NewExporter
// returns an exporter component.
func NewExporter() *Exporter {
	return (&Exporter{
		formatter: (&formatter{}).init(),
	}).init()
}

// /////////////////////////////////////////////////////////////////////////////
// Exporter: interface methods
// /////////////////////////////////////////////////////////////////////////////

// Processor
// returns a processor component.
func (o *Exporter) Processor() process.Processor {
	return o.processor
}

// Push
// log to exporter to do write / publish operation.
func (o *Exporter) Push(spans ...cores.Span) {
	atomic.AddInt32(&o.processing, 1)
	defer atomic.AddInt32(&o.processing, -1)

	if buf, err := o.formatter.Thrift(spans...); err == nil {
		_ = o.Upload(buf)
	}
}

// Upload span to jaeger.
func (o *Exporter) Upload(buf *bytes.Buffer) (err error) {
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
		req.Header.Set("Authorization",
			fmt.Sprintf("Basic %s",
				base64.StdEncoding.EncodeToString([]byte(usr+":"+pwd)),
			),
		)
	}

	// Send request.
	err = fasthttp.Do(req, res)
	return
}

// /////////////////////////////////////////////////////////////////////////////
// Exporter: events
// /////////////////////////////////////////////////////////////////////////////

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
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (o *Exporter) onPanic(_ context.Context, _ interface{}) {
	cores.Registry.Debugger("%s panic: %v", o.processor.Name())
}

// /////////////////////////////////////////////////////////////////////////////
// Exporter: init and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *Exporter) init() *Exporter {
	o.processor = process.New("jaeger-tracer-exporter").
		After(o.onAfter).
		Before(o.onBefore).
		Callback(o.onCall)
	return o
}

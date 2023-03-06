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

// Package tracer_term
// 打印到终端/控制台.
package tracer_term

import (
	"context"
	"fmt"
	"github.com/fuyibing/log/v5/common"
	"github.com/fuyibing/log/v5/configurer"
	"github.com/fuyibing/log/v5/tracers"
	"github.com/fuyibing/util/v8/process"
	"os"
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

func (o *executor) onAfter(_ context.Context) (ignored bool) { return }

func (o *executor) onBefore(_ context.Context) (ignored bool) { return }

func (o *executor) onCall(ctx context.Context) (ignored bool) {
	common.InternalInfo("<%s> signal listening", o.name)

	for {
		select {
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
	o.name = "tracer.term"
	o.processor = process.New(o.name).
		After(o.onAfter).
		Before(o.onBefore).
		Callback(o.onCall).
		Panic(o.onPanic)

	return o
}

func (o *executor) publish(spans ...tracers.Span) (err error) {
	var text string

	// 解析正文.
	if text, err = o.formatter.String(spans...); err != nil {
		return
	}

	// 打印日志.
	_, err = fmt.Fprintf(os.Stdout, "%s\n", text)
	return
}

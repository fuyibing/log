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
	"context"
	"encoding/binary"
	"fmt"
	"github.com/fuyibing/log/v5/configurer"
	"github.com/fuyibing/log/v5/loggers"
	"github.com/fuyibing/log/v5/tracers"
	"github.com/fuyibing/log/v5/tracers/tracer_jaeger/jaeger"
	"github.com/fuyibing/log/v5/tracers/tracer_jaeger/thrift"
	"strconv"
)

type (
	formatter struct{}
)

func (o *formatter) Byte(vs ...tracers.Span) ([]byte, error)           { return o.thrift(vs...) }
func (o *formatter) String(_ ...tracers.Span) (text string, err error) { return }

// /////////////////////////////////////////////////////////////////////////////
// Access
// /////////////////////////////////////////////////////////////////////////////

func (o *formatter) build(list ...tracers.Span) (batch *jaeger.Batch) {
	return &jaeger.Batch{
		Process: o.buildProcess(),
		Spans:   o.buildSpans(list...),
	}
}

func (o *formatter) buildLogs(list []loggers.Log) []*jaeger.Log {
	logs := make([]*jaeger.Log, 0)

	for _, x := range list {
		logs = append(logs, &jaeger.Log{
			Timestamp: x.Time().UnixMicro(),
			Fields: o.buildTagsMapper(x.Kv(), loggers.Kv{
				x.Time().Format("15:04:05.999999"): x.Text(),
				"log-level":                        x.Level(),
			}),
		})
	}
	return logs
}

func (o *formatter) buildProcess() *jaeger.Process {
	return &jaeger.Process{
		ServiceName: configurer.Config.GetTracerTopic(),
		Tags:        o.buildTagsMapper(tracers.Operator.GetResource()),
	}
}

func (o *formatter) buildSpan(sp tracers.Span) *jaeger.Span {
	var (
		tid  = sp.Trace().TraceId()
		pid  = sp.ParentSpanId()
		sid  = sp.SpanId()
		span = jaeger.NewSpan()
	)

	// Identify info.
	span.TraceIdHigh = int64(binary.BigEndian.Uint64(tid[0:8]))
	span.TraceIdLow = int64(binary.BigEndian.Uint64(tid[8:16]))
	span.SpanId = int64(binary.BigEndian.Uint64(sid[:]))
	span.ParentSpanId = int64(binary.BigEndian.Uint64(pid[:]))

	// Basic info and flags.
	span.OperationName = sp.Name()
	span.StartTime = sp.StartTime().UnixMicro()
	span.Duration = sp.Duration().Microseconds()
	span.Flags = 1

	// Extensions.
	span.Tags = o.buildTagsMapper(sp.Kv())
	span.Logs = o.buildLogs(sp.Logs())
	span.References = o.buildReference()
	return span
}

func (o *formatter) buildSpans(sps ...tracers.Span) []*jaeger.Span {
	list := make([]*jaeger.Span, 0)
	for _, sp := range sps {
		list = append(list, o.buildSpan(sp))
	}
	return list
}

func (o *formatter) buildReference() (refs []*jaeger.SpanRef) { return nil }

func (o *formatter) buildTagsMapper(attrs ...loggers.Kv) []*jaeger.Tag {
	var (
		tags = make([]*jaeger.Tag, 0)
	)

	for _, attr := range attrs {
		for k, v := range attr {
			tag := &jaeger.Tag{Key: k}

			switch v.(type) {
			case bool:
				val := v.(bool)
				tag.VType = jaeger.TagType_BOOL
				tag.VBool = &val
			case float32, float64:
				val, _ := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
				tag.VType = jaeger.TagType_DOUBLE
				tag.VDouble = &val
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				val, _ := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
				tag.VType = jaeger.TagType_LONG
				tag.VLong = &val
			case string:
				val := v.(string)
				tag.VType = jaeger.TagType_STRING
				tag.VStr = &val
			default:
				val := fmt.Sprintf("%v", v)
				tag.VType = jaeger.TagType_STRING
				tag.VStr = &val
			}

			tags = append(tags, tag)
		}
	}

	// Return
	// built tags.
	if len(tags) > 0 {
		return tags
	}
	return nil
}

func (o *formatter) init() *formatter { return o }

func (o *formatter) thrift(list ...tracers.Span) (buf []byte, err error) {
	var (
		bat = o.build(list...)
		ctx = context.Background()
		mem = thrift.NewTMemoryBuffer()
	)

	if err = bat.Write(ctx, thrift.NewTBinaryProtocolConf(mem, &thrift.TConfiguration{})); err == nil {
		buf = mem.Buffer.Bytes()
	}
	return
}

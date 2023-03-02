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
	"encoding/binary"
	"fmt"
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/log/v5/exporters/tracer_jaeger/jaeger"
	"github.com/fuyibing/log/v5/exporters/tracer_jaeger/thrift"
	"strconv"
)

type (
	Formatter interface {
		Generate(sps ...base.Span) (batch *jaeger.Batch)
		Thrift(list ...base.Span) (buf *bytes.Buffer, err error)
	}

	formatter struct{}
)

func (o *formatter) Generate(list ...base.Span) (batch *jaeger.Batch) {
	return &jaeger.Batch{
		Process: o.buildProcess(),
		Spans:   o.buildSpans(list...),
	}
}

func (o *formatter) Thrift(list ...base.Span) (buf *bytes.Buffer, err error) {
	var (
		bat = o.Generate(list...)
		ctx = context.Background()
		mem = thrift.NewTMemoryBuffer()
	)

	if err = bat.Write(ctx, thrift.NewTBinaryProtocolConf(mem, &thrift.TConfiguration{})); err != nil {
		return
	}

	buf = mem.Buffer
	return
}

// /////////////////////////////////////////////////////////////////////////////
// Formatter: access
// /////////////////////////////////////////////////////////////////////////////

func (o *formatter) buildLogs(list []base.Log) []*jaeger.Log {
	logs := make([]*jaeger.Log, 0)

	for _, x := range list {
		x.GetAttribute().
			Add(x.GetTime().Format("15:04:05.999999"), x.GetText()).
			Add("log-level", x.GetLevel()).
			Add("log-time", x.GetTime())

		logs = append(logs, &jaeger.Log{
			Timestamp: x.GetTime().UnixMicro(),
			Fields:    o.buildTagsMapper(x.GetAttribute()),
		})
	}
	return logs
}

func (o *formatter) buildProcess() *jaeger.Process {
	return &jaeger.Process{
		ServiceName: conf.Config.GetTracerTopic(),
		Tags:        o.buildTagsMapper(base.Resource),
	}
}

func (o *formatter) buildSpan(sp base.Span) *jaeger.Span {
	var (
		tid  = sp.GetTrace().GetTraceId()
		pid  = sp.GetParentSpanId()
		sid  = sp.GetSpanId()
		span = jaeger.NewSpan()
	)

	// Identify info.
	span.TraceIdHigh = int64(binary.BigEndian.Uint64(tid[0:8]))
	span.TraceIdLow = int64(binary.BigEndian.Uint64(tid[8:16]))
	span.SpanId = int64(binary.BigEndian.Uint64(sid[:]))
	span.ParentSpanId = int64(binary.BigEndian.Uint64(pid[:]))

	// Basic info and flags.
	span.OperationName = sp.GetName()
	span.StartTime = sp.GetStartTime().UnixMicro()
	span.Duration = sp.GetDuration().Microseconds()
	span.Flags = 1

	// Extensions.
	span.Tags = o.buildTagsMapper(sp.GetAttribute())
	span.Logs = o.buildLogs(sp.GetLogs())
	span.References = o.buildReference()
	return span
}

func (o *formatter) buildSpans(sps ...base.Span) []*jaeger.Span {
	list := make([]*jaeger.Span, 0)
	for _, sp := range sps {
		list = append(list, o.buildSpan(sp))
	}
	return list
}

func (o *formatter) buildReference() (refs []*jaeger.SpanRef) {
	return nil
}

func (o *formatter) buildTagsMapper(attrs ...base.Attribute) []*jaeger.Tag {
	var (
		tags = make([]*jaeger.Tag, 0)
	)

	for _, attr := range attrs {
		// Range attributes.
		for k, v := range attr {
			tag := &jaeger.Tag{Key: k}

			// Type verify.
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

func (o *formatter) init() *formatter {
	return o
}

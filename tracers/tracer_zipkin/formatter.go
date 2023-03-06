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

package tracer_zipkin

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/fuyibing/log/v5/configurer"
	"github.com/fuyibing/log/v5/loggers"
	"github.com/fuyibing/log/v5/tracers"
	"github.com/fuyibing/log/v5/tracers/tracer_zipkin/model"
)

type formatter struct{}

func (o *formatter) String(_ ...tracers.Span) (string, error) { return "", nil }

// Byte
// 转成链路数据.
func (o *formatter) Byte(vs ...tracers.Span) (body []byte, err error) {
	list := make([]*model.SpanModel, 0)
	for _, v := range vs {
		list = append(list, o.format(v))
	}
	return json.Marshal(list)
}

// format
// 数据格式化.
func (o *formatter) format(v tracers.Span) (sm *model.SpanModel) {
	sid := v.SpanId()
	pid := v.ParentSpanId()
	ptr := model.ID(binary.BigEndian.Uint64(pid[:]))
	tid := v.Trace().TraceId()

	// 标准结果.
	sm = &model.SpanModel{
		SpanContext: model.SpanContext{
			ID:       model.ID(binary.BigEndian.Uint64(sid[:])),
			ParentID: &ptr,
			TraceID:  model.TraceID{High: binary.BigEndian.Uint64(tid[:8]), Low: binary.BigEndian.Uint64(tid[8:])},
		},
		Name: v.Name(), Kind: model.Client,
		Timestamp: v.StartTime(), Duration: v.Duration(),
		LocalEndpoint: &model.Endpoint{ServiceName: configurer.Config.GetTracerTopic()},
	}

	// 日志
	sm.Annotations = o.genLogs(v.Logs()...)

	// 标签
	sm.Tags = o.genTags(tracers.Operator.GetResource(), v.Kv())
	return
}

// genLogs
// 生成日志.
func (o *formatter) genLogs(vs ...loggers.Log) []model.Annotation {
	if len(vs) == 0 {
		return nil
	}

	var (
		list = make([]model.Annotation, 0)
		text string
	)

	// 遍历日志.
	for _, v := range vs {
		text = fmt.Sprintf("[%s]", v.Level())
		// 键值.
		if kv := v.Kv(); len(kv) > 0 {
			text += fmt.Sprintf(" %v", kv.String())
		}
		// 正文.
		text += fmt.Sprintf(" %s", v.Text())
		// 堆栈.
		if v.Stack() {
			for _, item := range v.Stacks() {
				if item.Internal {
					continue
				}
				text += fmt.Sprintf("\n%s:%d", item.File, item.Line)
			}
		}
		// 清单.
		list = append(list, model.Annotation{Timestamp: v.Time(), Value: text})
	}
	return list
}

// genTags
// 生成标签.
func (o *formatter) genTags(kvs ...loggers.Kv) (mapper map[string]string) {
	mapper = make(map[string]string)
	for _, kv := range kvs {
		for k, v := range kv {
			mapper[k] = fmt.Sprintf("%v", v)
		}
	}
	return
}

func (o *formatter) init() *formatter { return o }

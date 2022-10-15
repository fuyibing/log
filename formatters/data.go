// author: wsfuyibing <websearch@163.com>
// date: 2022-10-14

package formatters

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/fuyibing/log/v3/base"
)

var dataPool *sync.Pool

// 默认JSON数据结构.
type data struct {
	// 1. 服务信息.

	Module     string `json:"module"`
	Pid        int    `json:"pid"`
	ServerAddr string `json:"serverAddr"`
	TaskId     int    `json:"taskId"`
	TaskName   string `json:"taskName"`

	// 2. 日志信息.

	Content  string  `json:"content"`
	Duration float64 `json:"duration"`
	Level    string  `json:"level"`
	Time     string  `json:"time"`

	// 3. 链路信息.

	Action        string `json:"action"`
	TraceId       string `json:"traceId"`
	ParentSpanId  string `json:"parentSpanId"`
	SpanId        string `json:"spanId"`
	Version       string `json:"version"`
	RequestId     string `json:"requestId"`
	RequestMethod string `json:"requestMethod"`
	RequestUrl    string `json:"requestUrl"`
}

// NewData
// 数据实例.
func NewData(line *base.Line, err error) *data {
	return dataPool.Get().(*data).
		before(line, err)
}

// String
// 转为JSON字符串.
func (o *data) String() (str string) {
	// 1. 监听结束.
	//    数据处理完成后, 释放实例回池.
	defer func() {
		o.after()
		dataPool.Put(o)
	}()

	// 2. 转换数据.
	if buf, err := json.Marshal(o); err == nil {
		str = string(buf)
		return
	}

	// 3. 空数据.
	str = "{}"
	return
}

func (o *data) after() *data {
	// 服务信息.

	o.Module = ""
	o.Pid = 0
	o.ServerAddr = ""

	// 日志信息.

	o.Content = ""
	o.Duration = 0
	o.Level = ""
	o.Time = ""

	// 链路信息.

	o.Action = ""
	o.TraceId = ""
	o.ParentSpanId = ""
	o.SpanId = ""
	o.Version = ""
	o.RequestId = ""
	o.RequestMethod = ""
	o.RequestUrl = ""
	return o
}

func (o *data) before(line *base.Line, err error) *data {
	// 服务信息.

	o.Module = base.LogName
	o.Pid = base.LogPid
	o.ServerAddr = fmt.Sprintf("%s:%d", base.LogHost, base.LogPort)

	// 日志信息.

	o.Content = line.Content
	o.Duration = line.Duration
	o.Level = line.Level.Name()
	o.Time = line.Time.Format(base.LogTimeFormat)

	// 链路信息.

	o.TraceId = line.TraceId
	o.ParentSpanId = line.ParentSpanId
	o.SpanId = line.SpanId
	o.Version = fmt.Sprintf("%s.%d", line.SpanPrefix, line.SpanOffset)
	o.RequestId = line.TraceId
	o.RequestMethod = line.RequestMethod
	o.RequestUrl = line.RequestUrl

	// 追加错误.
	if err != nil {
		o.Content += fmt.Sprintf(" << interrupt: %s", err.Error())
	}

	return o
}

func (o *data) init() *data {
	return o
}

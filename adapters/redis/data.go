// author: wsfuyibing <websearch@163.com>
// date: 2022-10-14

package redis

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/fuyibing/log/v3/adapters"
)

var (
	dataId   uint64 = 0
	dataPool *sync.Pool
)

type Data struct {
	// 服务信息.

	Module     string `json:"module"`
	Pid        int    `json:"pid"`
	ServerAddr string `json:"serverAddr"`
	TaskId     int    `json:"taskId"`
	TaskName   string `json:"taskName"`

	// 日志信息.

	Content  string  `json:"content"`
	Duration float64 `json:"duration"`
	Level    string  `json:"level"`
	Time     string  `json:"time"`

	// 链路信息.

	Action        string `json:"action"`
	TraceId       string `json:"traceId"`
	ParentSpanId  string `json:"parentSpanId"`
	SpanId        string `json:"spanId"`
	Version       string `json:"version"`
	RequestId     string `json:"requestId"`
	RequestMethod string `json:"requestMethod"`
	RequestUrl    string `json:"requestUrl"`

	index uint64
}

func NewData(x *adapters.Line) *Data {
	return dataPool.Get().(*Data).before(x)
}

func (o *Data) Key() string {
	return fmt.Sprintf("%s:%s:%d", Config.KeyPrefix, adapters.NodeId, o.index)
}

func (o *Data) String() (str string) {
	// 1. 监听结束.
	//    数据处理完成后, 释放实例回池.
	defer func() {
		o.after()
		dataPool.Put(o)
	}()

	// 2. 转换数据.
	if buf, err := json.Marshal(o); err == nil {
		str = string(buf)
	}
	return
}

func (o *Data) after() *Data {
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

func (o *Data) before(x *adapters.Line) *Data {
	o.index = x.GetIndex()

	// 服务信息.

	o.Module = x.Name
	o.Pid = x.Pid
	o.ServerAddr = fmt.Sprintf("%s:%d", x.Host, x.Port)

	// 日志信息.

	o.Content = x.Content
	o.Duration = x.Duration
	o.Level = x.Level.String()
	o.Time = x.Time.Format(x.TimeFormat)

	// 链路信息.

	o.Action = x.Action
	o.TraceId = x.TraceId
	o.ParentSpanId = x.ParentSpanId
	o.SpanId = x.SpanId
	o.Version = fmt.Sprintf("%s.%d", x.SpanPrefix, x.SpanOffset)
	o.RequestId = x.TraceId
	o.RequestMethod = x.RequestMethod
	o.RequestUrl = x.RequestUrl
	return o
}

func (o *Data) init() *Data {
	return o
}

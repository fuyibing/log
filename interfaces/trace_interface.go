// author: wsfuyibing <websearch@163.com>
// date: 2021-02-23

package interfaces

import (
	"net/http"
)

type TraceInterface interface {
	GenVersion(i int32) string
	GetSpanId() string
	GetSpanVersion() string
	GetTraceId() string
	IncrOffset() (before int32, after int32)
	RequestInfo() (method string, url string)
	UseDefault() TraceInterface
	UseRequest(req *http.Request) TraceInterface
}

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
// date: 2023-03-04

package tracers

import (
	"fmt"
	"github.com/fuyibing/log/v5/common"
	"github.com/fuyibing/log/v5/loggers"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	// Operator
	// 链路操作.
	Operator OperatorManager

	nilTraceId TraceId
	nilSpanId  SpanId
)

type (
	// OperatorManager
	// 链路操作接口.
	OperatorManager interface {
		// Generator
		// ID生成器.
		Generator() (generator *id)

		// GetExecutor
		// 执行器.
		GetExecutor() (executor Executor)

		// GetResource
		// 基础资源.
		GetResource() (kv loggers.Kv)

		// Push
		// 推送跨度.
		Push(span Span)

		// SetExecutor
		// 设置执行器.
		SetExecutor(executor Executor)
	}

	operator struct {
		executor  Executor
		generator *id
		name      string
		resource  loggers.Kv
	}
)

// /////////////////////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////////////////////

func (o *operator) Generator() (generator *id)       { return o.generator }
func (o *operator) GetExecutor() (executor Executor) { return o.executor }
func (o *operator) GetResource() (kv loggers.Kv)     { return o.resource }
func (o *operator) Push(span Span)                   { o.push(span) }
func (o *operator) SetExecutor(executor Executor)    { o.executor = executor }

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *operator) init() *operator {
	o.generator = (&id{}).init()
	o.name = "tracers.operator"
	o.resource = loggers.Kv{}

	o.initResource()
	return o
}

func (o *operator) initResource() {
	// 基础项.
	o.resource.Add("process.id", os.Getpid()).
		Add("system.arch", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)).
		Add("system.env", runtime.Version())

	// 主机名.
	if s, se := os.Hostname(); se == nil {
		o.resource.Add("system.host", s)
	}

	// IP地址.
	if l, le := net.InterfaceAddrs(); le == nil {
		ls := make([]string, 0)
		for _, la := range l {
			if ipn, ok := la.(*net.IPNet); ok && !ipn.IP.IsLoopback() {
				if ipn.IP.To4() != nil {
					ls = append(ls, ipn.IP.String())
				}
			}
		}
		o.resource.Add("system.addr", strings.Join(ls, ", "))
	}
}

func (o *operator) push(span Span) {
	// 链路禁用.
	if o.executor == nil {
		return
	}

	// 推送链路.
	if err := o.executor.Publish(span); err != nil {
		common.InternalFatal("<%s> send: %v", o.name, err)
	}
}

func init() { new(sync.Once).Do(func() { Operator = (&operator{}).init() }) }
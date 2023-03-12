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
	// for tracer operations.
	OperatorManager interface {
		// Generator
		// return id generator.
		Generator() (generator *id)

		// GetExecutor
		// return tracer executor.
		GetExecutor() (executor Executor)

		// GetResource
		// return operator key/value pairs.
		GetResource() (kv loggers.Kv)

		// Push
		// span component on to executor.
		Push(span Span)

		// SetExecutor
		// configure tracer executor.
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
	o.resource.Add("process.id", os.Getpid()).
		Add("system.arch", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)).
		Add("system.env", runtime.Version())

	if s, se := os.Hostname(); se == nil {
		o.resource.Add("system.host", s)
	}

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
	if o.executor == nil {
		return
	}
	if err := o.executor.Publish(span); err != nil {
		common.InternalFatal("<%s> send: %v", o.name, err)
	}
}

func init() { new(sync.Once).Do(func() { Operator = (&operator{}).init() }) }

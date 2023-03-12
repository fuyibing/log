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

package loggers

import (
	"github.com/fuyibing/log/v5/common"
	"github.com/fuyibing/log/v5/configurer"
	"sync"
)

var Operator OperatorManager

type (
	// OperatorManager
	// for logger operations.
	OperatorManager interface {
		// GetExecutor
		// return logger executor.
		GetExecutor() (executor Executor)

		// Push
		// log component on to executor.
		Push(kv Kv, level common.Level, format string, args ...interface{})

		// SetExecutor
		// configure logger executor.
		SetExecutor(executor Executor)
	}

	operator struct {
		executor Executor
		name     string
	}
)

// /////////////////////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////////////////////

func (o *operator) GetExecutor() Executor                                 { return o.executor }
func (o *operator) Push(k Kv, l common.Level, s string, a ...interface{}) { o.send(k, l, s, a...) }
func (o *operator) SetExecutor(v Executor)                                { o.executor = v }

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *operator) init() *operator {
	o.name = "loggers.operator"
	return o
}

func (o *operator) send(kv Kv, level common.Level, format string, args ...interface{}) {
	// Ignore
	// if executor not specified or log level is greater than configured.
	if o.executor == nil || !configurer.Config.LevelEnabled(level) {
		return
	}

	// Component
	// created for sender.
	v := NewLog(level, format, args...)

	// Copy
	// key/value pair.
	if kv != nil {
		v.SetKv(kv)
	}

	// Call specified executor
	// then push into it.
	if err := o.executor.Publish(v); err != nil {
		common.InternalInfo("<%s> publish: %v", o.name, err)
	}
}

func init() { new(sync.Once).Do(func() { Operator = (&operator{}).init() }) }

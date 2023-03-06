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

var (
	// Operator
	// 日志操作.
	Operator OperatorManager
)

type (
	// OperatorManager
	// 日志操作接口.
	OperatorManager interface {
		// GetExecutor
		// 执行器.
		GetExecutor() Executor

		// Push
		// 推送日志.
		Push(kv Kv, level common.Level, format string, args ...interface{})

		// SetExecutor
		// 设置执行器.
		SetExecutor(v Executor)
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

// 发送日志.
//
//   1. 日志级别
//   2. 绑定执行器
func (o *operator) send(kv Kv, level common.Level, text string, args ...interface{}) {
	// 日志禁用.
	// 1. 级别不匹配
	// 2. 执行器未定义
	if o.executor == nil || !configurer.Config.LevelEnabled(level) {
		return
	}

	// 建立日志.
	v := NewLog(level, text, args...)

	// 绑定KV.
	if kv != nil {
		v.SetKv(kv)
	}

	// 执行日志.
	// 发布到具体的执行器(如: Term/File/Kafka 等)存储.
	if err := o.executor.Publish(v); err != nil {
		common.InternalInfo("<%s> publish: %v", o.name, err)
	}
}

func init() { new(sync.Once).Do(func() { Operator = (&operator{}).init() }) }

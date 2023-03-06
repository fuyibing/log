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
	"fmt"
	"github.com/fuyibing/log/v5/common"
	"time"
)

type (
	// Log
	// 单条日志元素.
	Log interface {
		// Kv
		// Key/Value键值对.
		Kv() Kv

		// Level
		// 日志级别.
		Level() common.Level

		// Stack
		// 堆栈状态.
		Stack() bool

		// Stacks
		// 堆栈列表, 仅当级别为Fatal时有效.
		Stacks() []common.StackItem

		// Text
		// 日志正文.
		Text() string

		// Time
		// 记录时间.
		Time() time.Time

		// SetKv
		// 设置Key/Value键值对.
		SetKv(s Kv) Log
	}

	log struct {
		kv     Kv
		level  common.Level
		stack  bool
		stacks []common.StackItem
		text   string
		time   time.Time
	}
)

func NewLog(level common.Level, text string, args ...interface{}) Log {
	return (&log{
		level: level, text: fmt.Sprintf(text, args...),
		time: time.Now(),
	}).init()
}

// /////////////////////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////////////////////

func (o *log) Kv() Kv                     { return o.kv }
func (o *log) Level() common.Level        { return o.level }
func (o *log) Stack() bool                { return o.stack }
func (o *log) Stacks() []common.StackItem { return o.stacks }
func (o *log) Text() string               { return o.text }
func (o *log) Time() time.Time            { return o.time }

func (o *log) SetKv(s Kv) Log {
	if o.kv == nil {
		o.kv = Kv{}
	}

	o.kv.Copy(s)
	return o
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *log) init() *log {
	if o.stack = o.level == common.Fatal; o.stack {
		o.stacks = common.Backstack().Items
	}
	return o
}

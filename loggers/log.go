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
	// component for logger, stored with mixed.
	Log interface {
		Kv() Kv
		Level() common.Level
		SetKv(s Kv) Log
		Stack() bool
		Stacks() []common.StackItem
		Text() string
		Time() time.Time
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

func NewLog(level common.Level, format string, args ...interface{}) Log {
	return (&log{
		level: level, text: fmt.Sprintf(format, args...),
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

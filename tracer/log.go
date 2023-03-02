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
// date: 2023-03-01

package tracer

import (
	"fmt"
	"github.com/fuyibing/log/v5/base"
	"time"
)

type (
	// Log 用户日志.
	Log struct {
		Attribute base.Attribute
		L         base.Level
		T         time.Time
		Text      string
	}
)

func NewLog(level base.Level, text string, args ...interface{}) *Log {
	return &Log{
		Attribute: base.Attribute{},
		L:         level, T: time.Now(),
		Text: fmt.Sprintf(text, args...),
	}
}

// GetAttribute 日志属性.
func (o *Log) GetAttribute() base.Attribute { return o.Attribute }

// GetLevel 日志级别.
func (o *Log) GetLevel() base.Level { return o.L }

// GetTime 记录时间.
func (o *Log) GetTime() time.Time { return o.T }

// GetText 日志正文.
func (o *Log) GetText() string { return o.Text }

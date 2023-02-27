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
// date: 2023-02-26

package cores

import (
	"fmt"
	"github.com/fuyibing/log/v5/base"
	"time"
)

type (
	// Line
	// 用户日志.
	Line interface {
		// Add
		// 添加KV键值对到日志属性.
		Add(key string, value interface{}) Line

		// GetAttr
		// 获取日志的KV属性.
		GetAttr() Attr

		// GetLevel
		// 获取日志级别.
		GetLevel() base.Level

		// GetText
		// 获取日志正文.
		GetText() string

		// GetTime
		// 获取日志记录时间.
		GetTime() time.Time
	}

	line struct {
		attr  Attr
		level base.Level
		text  string
		time  time.Time
	}
)

// NewLine
// 创建用户日志实例.
func NewLine(level base.Level, text string, args ...interface{}) Line {
	return &line{
		attr:  NewAttr(),
		level: level,
		text:  fmt.Sprintf(text, args...),
		time:  time.Now(),
	}
}

// Add
// 添加KV键值对到日志属性.
func (o *line) Add(key string, value interface{}) Line {
	o.attr.Add(key, value)
	return o
}

// GetAttr
// 获取日志的KV属性.
func (o *line) GetAttr() Attr { return o.attr }

// GetLevel
// 获取日志级别.
func (o *line) GetLevel() base.Level { return o.level }

// GetText
// 获取日志正文.
func (o *line) GetText() string { return o.text }

// GetTime
// 获取日志记录时间.
func (o *line) GetTime() time.Time { return o.time }

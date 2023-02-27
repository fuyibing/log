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
	"sync"
)

type (
	// SpanLogs
	// 用于Span的Log组件.
	SpanLogs interface {
		// Add
		// 添加Log列表.
		Add(lines ...Line) SpanLogs

		// GetLines
		// 读取Log列表.
		GetLines() []Line
	}

	spanLogs struct {
		sync.RWMutex
		lines []Line
	}
)

// NewSpanLogs
// 创建SpanLogs组件.
func NewSpanLogs() SpanLogs { return &spanLogs{lines: make([]Line, 0)} }

// Add
// 添加Log列表.
func (o *spanLogs) Add(lines ...Line) SpanLogs {
	o.Lock()
	defer o.Unlock()

	o.lines = append(o.lines, lines...)
	return o
}

// GetLines
// 读取Log列表.
func (o *spanLogs) GetLines() []Line {
	o.RLock()
	defer o.RUnlock()

	return o.lines
}

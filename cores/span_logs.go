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
		// lines on SpanLogs.
		Add(lines ...Line) SpanLogs

		// GetLines
		// return log line list of the SpanLogs component.
		GetLines() []Line
	}

	spanLogs struct {
		sync.RWMutex
		lines []Line
	}
)

// NewSpanLogs
// returns a SpanLogs component.
func NewSpanLogs() SpanLogs {
	return &spanLogs{
		lines: make([]Line, 0),
	}
}

// Add
// lines on SpanLogs.
func (o *spanLogs) Add(lines ...Line) SpanLogs {
	o.Lock()
	defer o.Unlock()
	o.lines = append(o.lines, lines...)
	return o
}

// GetLines
// return log line list of the SpanLogs component.
func (o *spanLogs) GetLines() []Line {
	o.RLock()
	defer o.RUnlock()

	return o.lines
}

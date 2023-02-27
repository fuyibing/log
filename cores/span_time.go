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
	"time"
)

type (
	// SpanTime
	// 用于Span的时间组件.
	SpanTime interface {
		// End
		// 结束SpanTime组件.
		End() SpanTime

		// GetDuration
		// 获取SpanTime耗时.
		GetDuration() time.Duration

		// GetEnd
		// 获取SpanTime结束时间.
		GetEnd() time.Time

		// GetStart
		// 获取SpanTime开始时间.
		GetStart() time.Time
	}

	spanTime struct {
		sync.RWMutex
		start, end time.Time
	}
)

// NewSpanTime
// 创建SpanTime时间组件
func NewSpanTime() SpanTime { return &spanTime{start: time.Now()} }

// End
// 结束SpanTime组件.
func (o *spanTime) End() SpanTime {
	o.Lock()
	defer o.Unlock()

	o.end = time.Now()
	return o
}

// GetDuration
// 获取SpanTime耗时.
func (o *spanTime) GetDuration() time.Duration {
	o.RLock()
	defer o.RUnlock()

	return o.end.Sub(o.start)
}

// GetEnd
// 获取SpanTime结束时间.
func (o *spanTime) GetEnd() time.Time {
	o.RLock()
	defer o.RUnlock()

	return o.end
}

// GetStart
// 获取SpanTime开始时间.
func (o *spanTime) GetStart() time.Time {
	o.RLock()
	defer o.RUnlock()

	return o.start
}

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
		// 结束Span.
		End() SpanTime

		// GetDuration
		// 获取Span耗时.
		GetDuration() time.Duration

		// GetEnd
		// 获取Span结束时间.
		GetEnd() time.Time

		// GetStart
		// 获取Span开始时间.
		GetStart() time.Time
	}

	spanTime struct {
		sync.RWMutex
		start, end time.Time
	}
)

// NewSpanTime
// returns a SpanTime component.
func NewSpanTime() SpanTime {
	return &spanTime{
		start: time.Now(),
	}
}

// End
// set end time of the Span.
func (o *spanTime) End() SpanTime {
	o.Lock()
	o.end = time.Now()
	o.Unlock()
	return o
}

// GetDuration
// returns a time.Duration of the Span lifetime.
func (o *spanTime) GetDuration() time.Duration {
	o.RLock()
	defer o.RUnlock()
	return o.end.Sub(o.start)
}

// GetEnd
// returns a time.Time of the Span ended.
func (o *spanTime) GetEnd() time.Time {
	o.RLock()
	defer o.RUnlock()
	return o.end
}

// GetStart
// returns a time.Time of the Span started.
func (o *spanTime) GetStart() time.Time {
	o.RLock()
	defer o.RUnlock()
	return o.start
}

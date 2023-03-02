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
// date: 2023-03-02

package log

import (
	"context"
	"github.com/fuyibing/log/v5/tracer"
	"github.com/fuyibing/log/v5/traces"
	"net/http"
)

// NewTrace 生成跟踪组件.
func NewTrace(name string) traces.Trace {
	return tracer.NewTrace(name)
}

// NewTraceFromContext 生成跟踪组件.
func NewTraceFromContext(ctx context.Context, name string) traces.Trace {
	return tracer.NewTraceFromContext(ctx, name)
}

// NewTraceFromRequest 生成跟踪组件.
func NewTraceFromRequest(req *http.Request, name string) traces.Trace {
	return tracer.NewTraceFromRequest(req, name)
}

// Span 获取链路.
func Span(ctx context.Context) (span traces.Span, exists bool) {
	if v := ctx.Value(traces.ContextKeyForTrace); v != nil {
		if s, ok := v.(traces.Span); ok {
			return s, true
		}
	}
	return nil, false
}

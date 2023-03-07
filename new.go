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
// date: 2023-03-06

package log

import (
	"context"
	"github.com/fuyibing/log/v5/tracers"
	"net/http"
)

// NewSpan 返回 tracers.Span 组件. 此过程先创建 tracers.Trace 组件, 然后基于此组
// 件创建 tracers.Span 组件并返回.
func NewSpan(name string) (span tracers.Span) {
	return tracers.NewSpan(name)
}

// NewSpanFromContext 返回 tracers.Span 组件. 若 context.Context 绑定过
// tracers.Span 组件则基于此创建子 tracers.Span 并返回, 若绑定过 tracers.Trace
// 则基于此创建新的 tracers.Span 并返回, 反之则使用和 NewSpan 相同逻辑.
func NewSpanFromContext(ctx context.Context, name string) (span tracers.Span) {
	return tracers.NewSpanFromContext(ctx, name)
}

// NewSpanFromRequest 返回 tracers.Span 组件. 基于 HTTP 请求创建并返回, 创建
// 过程同 NewSpan 逻辑, 不同点在于此过程打通服务间链路.
func NewSpanFromRequest(req *http.Request, name string) (span tracers.Span) {
	return tracers.NewSpanFromRequest(req, name)
}

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

package log

import (
	"context"
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/cores"
	"net/http"
)

func NewTrace(name string) cores.Trace {
	return cores.NewTrace(name)
}

func NewTraceFromContext(ctx context.Context, name string) cores.Trace {
	return cores.NewTraceFromContext(ctx, name)
}

func NewTraceFromRequest(req *http.Request, name string) cores.Trace {
	return cores.NewTraceFromRequest(req, name)
}

// Trace
// 从上下文上读取链路, 若不存在则返回nil.
//
// 此方法为只读, 在实际开发过程中若需要基于上下文创建链路可使用
// NewTraceFromContext 方法替代.
func Trace(ctx context.Context) (trace cores.Trace, exists bool) {
	if v := ctx.Value(base.ContentKeyTrace); v != nil {
		trace, exists = v.(cores.Trace)
	}
	return
}

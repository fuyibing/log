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
	return Manager.NewTrace(name)
}

func NewTraceFromContext(ctx context.Context, name string) cores.Trace {
	return Manager.NewTraceFromContext(ctx, name)
}

func NewTraceFromRequest(req *http.Request, name string) cores.Trace {
	return Manager.NewTraceFromRequest(req, name)
}

// Trace
// returns a core.Trace component from context.Context stored value.
func Trace(ctx context.Context) (trace cores.Trace, exists bool) {
	if v := ctx.Value(base.ContentKeyTrace); v != nil {
		trace, exists = v.(cores.Trace)
	}
	return
}

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

// NewSpan
// 链路跟踪.
func NewSpan(name string) tracers.Span {
	return tracers.NewSpan(name)
}

// NewSpanFromContext
// 链路跟踪.
func NewSpanFromContext(ctx context.Context, name string) tracers.Span {
	return tracers.NewSpanWithContext(ctx, name)
}

// NewSpanFromRequest
// 链路跟踪.
func NewSpanFromRequest(req *http.Request, name string) tracers.Span {
	return tracers.NewSpanWithRequest(req, name)
}

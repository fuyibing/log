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

// NewSpan returns a tracers.Span component.
func NewSpan(name string) (span tracers.Span) {
	return tracers.NewSpan(name)
}

// NewSpanFromContext returns a tracers.Span component, based on specified
// context.Context.
func NewSpanFromContext(ctx context.Context, name string) (span tracers.Span) {
	return tracers.NewSpanFromContext(ctx, name)
}

// NewSpanFromRequest returns a tracers.Span component, based on http request
// and context.Context.
func NewSpanFromRequest(req *http.Request, name string) (span tracers.Span) {
	return tracers.NewSpanFromRequest(req, name)
}

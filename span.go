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
// date: 2023-03-03

package log

import (
	"context"
	"github.com/fuyibing/log/v5/tracers"
)

// Span returns a tracers.Span component if valued in context, otherwise nil
// returned.
func Span(ctx context.Context) (span tracers.Span, exists bool) {
	if v, ok := ctx.Value(tracers.ContextKey).(tracers.Span); ok {
		return v, true
	}

	return nil, false
}

// Trace returns a tracers.Trace component if valued in context, otherwise nil
// returned.
func Trace(ctx context.Context) (trace tracers.Trace, exists bool) {
	if g := ctx.Value(tracers.ContextKey); g != nil {
		// Use span component.
		if v, ok := g.(tracers.Span); ok {
			return v.Trace(), true
		}

		// Use trace component.
		if v, ok := g.(tracers.Trace); ok {
			return v, true
		}
	}

	return nil, false
}

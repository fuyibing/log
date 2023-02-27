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
// date: 2023-02-27

package log

import (
	"context"
	"net/http"
)

var (
	ctx context.Context
	err error
	req *http.Request
)

func ExampleDebug() {
	Debug("debug message")
	// Output:
	// [2023-02-27 09:10:11.234567][DEBUG] debug message
}

func ExampleError() {
	Error("error message: %v", err)
	// Output:
	// [2023-02-27 09:10:11.234567][ERROR] error message: HTTP 500 Server internal error
}

func ExampleFatal() {
	Fatal("panic on: %v", err)
	// Output:
	// [2023-02-27 09:10:11.234567][FATAL] panic on: unknown dsn configuration
}

func ExampleField_Debug() {
	Field{}.
		Add("key", "value").
		Add("type", 1).
		Debug("message")
	// Output:
	// [2023-02-27 09:10:11.234567][DEBUG] {"key":"value","type":1} message
}

func ExampleField_Error() {
	Field{}.
		Add("key", "value").
		Add("type", 1).
		Error("message")
	// Output:
	// [2023-02-27 09:10:11.234567][ERROR] {"key":"value","type":1} message
}

func ExampleField_Fatal() {
	Field{}.
		Add("key", "value").
		Add("type", 1).
		Fatal("message")
	// Output:
	// [2023-02-27 09:10:11.234567][FATAL] {"key":"value","type":1} message
}

func ExampleField_Info() {
	Field{}.
		Add("key", "value").
		Add("type", 1).
		Info("message")
	// Output:
	// [2023-02-27 09:10:11.234567][INFO] {"key":"value","type":1} message
}

func ExampleField_Warn() {
	Field{}.
		Add("key", "value").
		Add("type", 1).
		Warn("message")
	// Output:
	// [2023-02-27 09:10:11.234567][WARN] {"key":"value","type":1} message
}

func ExampleInfo() {
	Info("my message")
	// Output:
	// [2023-02-27 09:10:11.234567][INFO] my message
}

func ExampleNewTrace() {
	trace := NewTrace("/health")

	span := trace.NewSpan("health check")
	defer span.End()

	span.Logger().Info("new trace created for root")
}

func ExampleNewTraceFromContext() {
	trace := NewTraceFromContext(ctx, "my trace")

	span := trace.NewSpan("my span")
	defer span.End()

	span.Logger().Debug("create trace if not valued in context")
}

func ExampleNewTraceFromRequest() {
	trace := NewTraceFromRequest(req, "my trace")

	span := trace.NewSpan("my span")
	defer span.End()

	span.Logger().Debug("create trace if not valued in http request")
}

func ExampleSpan() {
	if span, exists := Span(ctx); exists {
		span.Logger().Info("记录在 Span 上的日志内容")
	} else {
		Info("日志内容")
	}
}

func ExampleTrace() {
	if trace, exists := Trace(ctx); exists {
		span := trace.NewSpan("my span")
		defer span.End()

		span.Logger().Info("上游已定义了 Trace")
	} else {
		Info("上游 Trace 未定义")
	}
}

func ExampleWarn() {
	Warn("warning message")
	// Output:
	// [2023-02-27 09:10:11.234567][WARN] warning message
}

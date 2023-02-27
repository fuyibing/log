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

package base

const (
	ContextKeySpan  = "__OPEN_TELEMETRY_SPAN__"
	ContentKeyTrace = "__OPEN_TELEMETRY_TRACE__"

	DefaultLoggerExporter = "term"
	DefaultLoggerLevel    = "INFO"

	OpenTracingSample  = "X-B3-Sample"
	OpenTracingSpanId  = "X-B3-Spanid"
	OpenTracingTraceId = "X-B3-Traceid"

	ResourceArch        = "arch"
	ResourceEnvironment = "environment"
	ResourceHostAddress = "host.addr"
	ResourceHostName    = "host.name"
	ResourceProcessId   = "process.pid"

	ResourceServiceName    = "service.name"
	ResourceServicePort    = "service.port"
	ResourceServiceVersion = "service.version"

	ResourceHttpProtocol      = "http.protocol"
	ResourceHttpHeader        = "http.request.header"
	ResourceHttpRequestMethod = "http.request.method"
	ResourceHttpRequestUrl    = "http.request.url"
	ResourceHttpUserAgent     = "http.user.agent"
)

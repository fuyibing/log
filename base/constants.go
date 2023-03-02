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
// date: 2023-03-01

package base

const (
	ContextKeyForTrace = "__LOG_TRACE_CONTEXT_TRACE__"
)

const (
	BucketBatch       = 100
	BucketConcurrency = 10
	BucketCapacity    = 30000
	BucketFrequency   = 300

	OpenTracingSampled = "X-B3-Sample"
	OpenTracingSpanId  = "X-B3-Spanid"
	OpenTracingTraceId = "X-B3-Traceid"

	TracerTopic = "log-trace"

	ResourceArch              = "host.arch"
	ResourceEnvironment       = "host.env"
	ResourceHostAddress       = "host.addr"
	ResourceHostName          = "host.name"
	ResourceProcessId         = "process.pid"
	ResourceServiceName       = "service.name"
	ResourceServicePort       = "service.port"
	ResourceServiceVersion    = "service.version"
	ResourceHttpHeader        = "http.header"
	ResourceHttpProtocol      = "http.protocol"
	ResourceHttpRequestMethod = "http.request.method"
	ResourceHttpRequestUrl    = "http.request.url"
	ResourceHttpUserAgent     = "http.user.agent"
)

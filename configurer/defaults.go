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

package configurer

import (
	"github.com/fuyibing/log/v5/common"
)

const (
	defaultBucketBatch       = 100
	defaultBucketCapacity    = 30000
	defaultBucketConcurrency = 10
	defaultBucketFrequency   = 100
)

const (
	defaultLoggerExporter = "term"
	defaultLoggerLevel    = common.Info
	defaultTracerTopic    = "log-trace"
)

const (
	defaultOpenTracingSampled = "X-B3-Sampled"
	defaultOpenTracingSpanId  = "X-B3-Spanid"
	defaultOpenTracingTraceId = "X-B3-Traceid"
)

const (
	defaultFileLoggerExt    = "log"
	defaultFileLoggerFolder = "2006-01"
	defaultFileLoggerName   = "2006-01-02"
	defaultFileLoggerPath   = "./logs"
)

const (
	defaultFileTracerExt    = "trace.log"
	defaultFileTracerFolder = "06-01"
	defaultFileTracerName   = "06-01-02"
	defaultFileTracerPath   = "./logs"
)

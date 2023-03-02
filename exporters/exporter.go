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

package exporters

import (
	"fmt"
	"github.com/fuyibing/log/v5/traces"
	"net"
	"os"
	"runtime"
	"strings"
)

var (
	// Exporter
	// 上报实例.
	Exporter ExporterManager
)

type (
	// ExporterManager
	// 上报管理器.
	ExporterManager interface {
		GetLogger() traces.LoggerExporter
		GetTracer() traces.TracerExporter
		PutLogger(v traces.Log)
		PutTracer(v traces.Span)
		SetLogger(v traces.LoggerExporter)
		SetTracer(v traces.TracerExporter)
	}

	exporter struct {
		loggerEnabled  bool
		loggerExporter traces.LoggerExporter
		tracerEnabled  bool
		tracerExporter traces.TracerExporter
	}
)

func (o *exporter) GetLogger() traces.LoggerExporter {
	return o.loggerExporter
}

func (o *exporter) GetTracer() traces.TracerExporter {
	return o.tracerExporter
}

func (o *exporter) PutLogger(v traces.Log) {
	if o.loggerEnabled {
		if err := o.loggerExporter.Send(v); err != nil {
			traces.InternalError("<%s> send error: %v",
				o.loggerExporter.Processor().Name(),
				err,
			)
		}
	}
}

func (o *exporter) PutTracer(v traces.Span) {
	if o.tracerEnabled {
		if err := o.tracerExporter.Send(v); err != nil {
			traces.InternalError("<%s> send error: %v",
				o.tracerExporter.Processor().Name(),
				err,
			)
		}
	}
}

func (o *exporter) SetLogger(v traces.LoggerExporter) {
	o.loggerEnabled = v != nil
	o.loggerExporter = v
}

func (o *exporter) SetTracer(v traces.TracerExporter) {
	o.tracerEnabled = v != nil
	o.tracerExporter = v
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *exporter) init() *exporter {
	o.initResource()
	return o
}

func (o *exporter) initResource() {
	traces.Resource.Add(traces.ResourceProcessId, os.Getpid()).
		Add(traces.ResourceArch, fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)).
		Add(traces.ResourceEnvironment, runtime.Version())

	// 主机名.
	if s, se := os.Hostname(); se == nil {
		traces.Resource.Add(traces.ResourceHostName, s)
	}

	// 主机IP列表.
	// 例如: 172.16.0.103
	if l, le := net.InterfaceAddrs(); le == nil {
		ls := make([]string, 0)
		for _, la := range l {
			if ipn, ok := la.(*net.IPNet); ok && !ipn.IP.IsLoopback() {
				if ipn.IP.To4() != nil {
					ls = append(ls, ipn.IP.String())
				}
			}
		}
		traces.Resource.Add(traces.ResourceHostAddress, strings.Join(ls, ", "))
	}
}

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
	"github.com/fuyibing/log/v5/base"
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
		GetLogger() base.LoggerExporter
		GetTracer() base.TracerExporter
		PutLogger(v base.Log)
		PutTracer(v base.Span)
		SetLogger(v base.LoggerExporter)
		SetTracer(v base.TracerExporter)
	}

	exporter struct {
		loggerEnabled  bool
		loggerExporter base.LoggerExporter
		tracerEnabled  bool
		tracerExporter base.TracerExporter
	}
)

func (o *exporter) GetLogger() base.LoggerExporter {
	return o.loggerExporter
}

func (o *exporter) GetTracer() base.TracerExporter {
	return o.tracerExporter
}

func (o *exporter) PutLogger(v base.Log) {
	if o.loggerEnabled {
		if err := o.loggerExporter.Send(v); err != nil {
			base.InternalError("<%s> send error: %v",
				o.loggerExporter.Processor().Name(),
				err,
			)
		}
	}
}

func (o *exporter) PutTracer(v base.Span) {
	if o.tracerEnabled {
		if err := o.tracerExporter.Send(v); err != nil {
			base.InternalError("<%s> send error: %v",
				o.tracerExporter.Processor().Name(),
				err,
			)
		}
	}
}

func (o *exporter) SetLogger(v base.LoggerExporter) {
	o.loggerEnabled = v != nil
	o.loggerExporter = v
}

func (o *exporter) SetTracer(v base.TracerExporter) {
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
	base.Resource.Add(base.ResourceProcessId, os.Getpid()).
		Add(base.ResourceArch, fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)).
		Add(base.ResourceEnvironment, runtime.Version())

	// 主机名.
	if s, se := os.Hostname(); se == nil {
		base.Resource.Add(base.ResourceHostName, s)
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
		base.Resource.Add(base.ResourceHostAddress, strings.Join(ls, ", "))
	}
}

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

package cores

import (
	"fmt"
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/util/v8/process"
	"net"
	"os"
	"runtime"
	"strings"
)

var (
	Registry RegistryManager
)

type (
	// LoggerExporter
	// .
	LoggerExporter interface {
		Processor() process.Processor
		Push(lines ...Line)
	}

	// TracerExporter
	// .
	TracerExporter interface {
		Processor() process.Processor
		Push(spans ...Span)
	}

	RegistryManager interface {
		Debugger(text string, args ...interface{})
		LoggerEnabled() bool
		LoggerExporter() LoggerExporter
		LoggerPush(data interface{}, level base.Level, text string, args ...interface{})
		RegisterLoggerExporter(v LoggerExporter) RegistryManager
		RegisterTracerExporter(v TracerExporter) RegistryManager
		Resource() Attr
		TracerEnabled() bool
		TracerExporter() TracerExporter
		Update()
	}

	registry struct {
		loggerEnabled  bool
		loggerExporter LoggerExporter
		resource       Attr
		tracerEnabled  bool
		tracerExporter TracerExporter
	}
)

func (o *registry) Debugger(text string, args ...interface{}) {
	println(fmt.Sprintf(text, args...))
}

func (o *registry) LoggerEnabled() bool {
	return o.loggerEnabled
}

func (o *registry) LoggerExporter() LoggerExporter {
	return o.loggerExporter
}

func (o *registry) LoggerPush(data interface{}, level base.Level, text string, args ...interface{}) {
	// Ignore log if not enabled.
	if !o.loggerEnabled {
		return
	}

	// Init log line.
	log := NewLine(level, text, args...)
	log.GetAttr().With(data)

	// Push to exporter.
	o.loggerExporter.Push(log)

	// // Log fields.
	// if data != nil {
	// 	m, ok := data.(map[string]interface{})
	// 	fmt.Printf("---------- convert: %v\n", ok)
	// 	// println("convert: ", ok)
	// 	if ok {
	// 		for k, v := range m {
	// 			log.Add(k, v)
	// 		}
	// 	}
	// }
}

func (o *registry) TracerEnabled() bool {
	return o.tracerEnabled
}

func (o *registry) TracerExporter() TracerExporter {
	return o.tracerExporter
}

func (o *registry) RegisterLoggerExporter(v LoggerExporter) RegistryManager {
	o.loggerExporter = v
	o.loggerEnabled = v != nil
	return o
}

func (o *registry) RegisterTracerExporter(v TracerExporter) RegistryManager {
	o.tracerExporter = v
	o.tracerEnabled = v != nil
	return o
}

func (o *registry) Resource() Attr {
	return o.resource
}

func (o *registry) Update() {
	o.initResource()
	o.initResourceService()
}

func (o *registry) init() *registry {
	o.resource = NewAttr()
	return o
}

// initResource
// initialize resource key/value pairs.
func (o *registry) initResource() {
	// Process info.
	//
	// - runtime.pid: process id
	// - runtime.version: go version
	o.resource.Add(base.ResourceProcessId, os.Getpid())
	o.resource.Add(base.ResourceEnvironment, runtime.Version())

	// Host name.
	if s, se := os.Hostname(); se == nil {
		o.resource.Add(base.ResourceHostName, s)
	}

	// Server arch platform.
	o.resource.Add(base.ResourceArch, runtime.GOARCH)

	// Host addr, IPv4.
	if l, le := net.InterfaceAddrs(); le == nil {
		ls := make([]string, 0)
		for _, la := range l {
			if ipn, ok := la.(*net.IPNet); ok && !ipn.IP.IsLoopback() {
				if ipn.IP.To4() != nil {
					ls = append(ls, ipn.IP.String())
				}
			}
		}
		o.resource.Add(base.ResourceHostAddress, strings.Join(ls, ", "))
	}
}

func (o *registry) initResourceService() {
	if s := conf.Config.GetServiceName(); s != "" {
		o.resource.Add(base.ResourceServiceName, s)
	}
	if s := conf.Config.GetServicePort(); s > 0 {
		o.resource.Add(base.ResourceServicePort, s)
	}
	if s := conf.Config.GetServiceVersion(); s != "" {
		o.resource.Add(base.ResourceServiceVersion, s)
	}
}

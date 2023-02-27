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
	// Registry
	// 注册管理器.
	Registry RegistryManager
)

type (
	// LoggerExporter
	// Logger 输出接口.
	LoggerExporter interface {
		// Processor
		// 获取执行器.
		Processor() process.Processor

		// Push
		// 发布 Log 日志.
		Push(lines ...Line)
	}

	// TracerExporter
	// Tracer 输出接口.
	TracerExporter interface {
		// Processor
		// 获取执行器.
		Processor() process.Processor

		// Push
		// 发布 Span 跨度.
		Push(spans ...Span)
	}

	// RegistryManager
	// 注册管理器接口.
	RegistryManager interface {
		// Debugger
		// 打印调试信息.
		Debugger(text string, args ...interface{})

		// LoggerEnabled
		// 获取 Logger 启用状态.
		LoggerEnabled() bool

		// LoggerExporter
		// 获取 Logger 输出接口.
		LoggerExporter() LoggerExporter

		// LoggerPush
		// 发布日志到 Logger 输出接口.
		LoggerPush(data interface{}, level base.Level, text string, args ...interface{})

		// RegisterLoggerExporter
		// 注册 Logger 输出接口.
		RegisterLoggerExporter(v LoggerExporter) RegistryManager

		// RegisterTracerExporter
		// 注册 Tracer 输出接口.
		RegisterTracerExporter(v TracerExporter) RegistryManager

		// Resource
		// 获取资源属性.
		Resource() Attr

		// TracerEnabled
		// 获取 Tracer 启用状态.
		TracerEnabled() bool

		// TracerExporter
		// 获取 Tracer 输出接口.
		TracerExporter() TracerExporter

		// Update
		// 更新属性.
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

// Debugger
// 打印调试信息.
func (o *registry) Debugger(text string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, fmt.Sprintf(text, args...)+"\n")
}

// LoggerEnabled
// 获取 Logger 启用状态.
func (o *registry) LoggerEnabled() bool {
	return o.loggerEnabled
}

// LoggerExporter
// 获取 Logger 输出接口.
func (o *registry) LoggerExporter() LoggerExporter {
	return o.loggerExporter
}

// LoggerPush
// 发布日志到 Logger 输出接口.
func (o *registry) LoggerPush(data interface{}, level base.Level, text string, args ...interface{}) {
	// 忽略 Log.
	if !o.loggerEnabled {
		return
	}

	// 创建 Log 并转发.
	log := NewLine(level, text, args...)
	log.GetAttr().With(data)
	o.loggerExporter.Push(log)
}

// TracerEnabled
// 获取 Tracer 启用状态.
func (o *registry) TracerEnabled() bool {
	return o.tracerEnabled
}

// TracerExporter
// 获取 Tracer 输出接口.
func (o *registry) TracerExporter() TracerExporter {
	return o.tracerExporter
}

// RegisterLoggerExporter
// 注册 Logger 输出接口.
func (o *registry) RegisterLoggerExporter(v LoggerExporter) RegistryManager {
	o.loggerExporter = v
	o.loggerEnabled = v != nil
	return o
}

// RegisterTracerExporter
// 注册 Tracer 输出接口.
func (o *registry) RegisterTracerExporter(v TracerExporter) RegistryManager {
	o.tracerExporter = v
	o.tracerEnabled = v != nil
	return o
}

// Resource
// 获取资源属性.
func (o *registry) Resource() Attr {
	return o.resource
}

// Update
// 更新属性.
func (o *registry) Update() {
	o.initResource()
	o.initResourceService()
}

// init
// 注册中心构造.
func (o *registry) init() *registry {
	o.resource = NewAttr()
	return o
}

// initResource
// 初始化系统参数.
func (o *registry) initResource() {
	// 进程与环境.
	//
	// - runtime.pid: process id
	// - runtime.version: go version
	o.resource.Add(base.ResourceProcessId, os.Getpid())
	o.resource.Add(base.ResourceEnvironment, runtime.Version())

	// 主机名.
	if s, se := os.Hostname(); se == nil {
		o.resource.Add(base.ResourceHostName, s)
	}

	// 系统与内核
	// 例如: darwin amd64.
	o.resource.Add(base.ResourceArch, runtime.GOOS+" "+runtime.GOARCH)

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
		o.resource.Add(base.ResourceHostAddress, strings.Join(ls, ", "))
	}
}

// initResourceService
// 初始化服务参数.
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

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

package log

import (
	"context"
	"github.com/fuyibing/log/v5/built_in"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/log/v5/cores"
	"github.com/fuyibing/util/v8/process"
	"sync"
	"time"
)

var (
	// Manager
	// 管理器入口.
	Manager Management
)

type (
	// Management
	// 管理器接口.
	//
	// 本管理器在首次引入时自动读取配置文件(config/log.yaml)参数. 若未定义配置文件
	// 则需要在启动服务前通过代码配置.
	Management interface {
		// Start
		// 启动管理器.
		//
		// 调用在项目 main 的入口处, 全局仅定义一次, 此方法为阻塞模式, 需使用 go
		// 协程调用, 当启动后, 异步模式生效.
		Start(ctx context.Context) error

		// Stop
		// 退出管理器.
		//
		// 在服务退出前调用, 调用后可以确保异步处理的 Log/Trace 能正常完成, 避免
		// 丢失.
		Stop()
	}

	management struct {
		processor process.Processor
	}
)

// Start
// 启动管理器.
func (o *management) Start(ctx context.Context) error {
	return o.processor.Start(ctx)
}

// Stop
// 退出管理器.
func (o *management) Stop() {
	o.processor.Stop()

	// 等待完成.
	// 每 100 毫秒检查 Log/Trace 是否上报完成, 直到上报完成或30秒超时后退出.
	var (
		duration = time.Millisecond * 100
		limit    = 300
	)
	for i := 0; i < limit; i++ {
		if o.processor.Stopped() {
			break
		}
		time.Sleep(duration)
	}
}

// init
// 构造管理器.
func (o *management) init() *management {
	// 创建执行器.
	o.processor = process.New("log-tracer").
		Before(o.onBuiltinLogger, o.onBuiltinTracer, o.onBeforeDone).
		Callback(o.onListen).
		Panic(o.onPanic)
	return o
}

// onBeforeDone
// 启动前校验.
func (o *management) onBeforeDone(_ context.Context) (ignored bool) {
	cores.Registry.Update()
	return
}

// onBuiltinLogger
// 校验内置 Logger.
func (o *management) onBuiltinLogger(_ context.Context) (ignored bool) {
	if exporter := built_in.BuiltinLogger(conf.Config.GetLoggerExporter()).Exporter(); exporter != nil {
		cores.Registry.RegisterLoggerExporter(exporter)
		o.processor.Add(exporter.Processor())
	}
	return
}

// onBuiltinTracer
// 校验内置 Tracer.
func (o *management) onBuiltinTracer(_ context.Context) (ignored bool) {
	if exporter := built_in.BuiltinTracer(conf.Config.GetTracerExporter()).Exporter(); exporter != nil {
		cores.Registry.RegisterTracerExporter(exporter)
		o.processor.Add(exporter.Processor())
	}
	return
}

// onListen
// 阻塞并监听上下文信号.
func (o *management) onListen(ctx context.Context) (ignored bool) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

// onPanic
// 管理器运行异常.
func (o *management) onPanic(_ context.Context, v interface{}) {
	cores.Registry.Debugger("log.Manager fatal: %v", v)
}

func init() {
	new(sync.Once).Do(func() {
		Manager = (&management{}).init()
	})
}

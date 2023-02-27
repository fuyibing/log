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

package main

import (
	"github.com/fuyibing/log/v5"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/log/v5/cores"
	"github.com/fuyibing/log/v5/exporters/logger_term"
)

func init() {
	conf.Config.With(
		conf.TracerExporter("kafka"),
		conf.TracerTopic("log-trace"),

		conf.ServiceName("my-app"),
		conf.ServiceVersion("1.2.3"),
		conf.ServicePort(3721),

		conf.JaegerTracerContentType("application/x-thrift"),
		conf.JaegerTracerEndpoint("http://localhost:14268/api/traces"),
	)

	conf.LogExporter("kafka")
	conf.LogLevel("debug")

	cores.Registry.RegisterLoggerExporter(logger_term.New())
	// cores.Registry.RegisterTracerExporter(tracer_term.NewExporter())
	// cores.Registry.RegisterTracerExporter(tracer_jaeger.NewExporter())
	cores.Registry.Update()
}

func main() {
	log.Info("log.Info()")

	log.Field{}.
		Add("key", "value").
		Info("log.Field.Info()")

	log.Field{}.
		Add("key", "value").
		Info("log.Field.Info()")

	spa := log.NewTrace("trace").NewSpan("span")
	defer spa.End()

	spa.Logger().
		Add("key 1", "value 1").
		Add("key 2", "value 2").
		Debug("debug in span")

	spa.Logger().
		Info("log.cores.SpanLogger.Inf() info in span")

	// println("span name: ", sp.GetName())
}

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
// date: 2023-03-05

package main

import (
	"context"
	"github.com/fuyibing/log/v5"
	"github.com/fuyibing/log/v5/loggers/logger_file"
	"github.com/fuyibing/log/v5/loggers/logger_kafka"
	"github.com/fuyibing/log/v5/loggers/logger_term"
	"github.com/fuyibing/log/v5/tracers"
	"github.com/fuyibing/log/v5/tracers/tracer_file"
	"github.com/fuyibing/log/v5/tracers/tracer_jaeger"
	"github.com/fuyibing/log/v5/tracers/tracer_term"
	"github.com/fuyibing/log/v5/tracers/tracer_zipkin"
	"time"
)

var (
	ctx = context.Background()
)

func init() {
	// initLoggerFile()
	// initLoggerKafka()
	// initLoggerTerm()

	// initTracerFile()
	// initTracerJaeger()
	// initTracerTerm()
	// initTracerZipkin()
}

func initLoggerFile()  { log.Manager.Logger().SetExecutor(logger_file.New()) }
func initLoggerKafka() { log.Manager.Logger().SetExecutor(logger_kafka.New()) }
func initLoggerTerm()  { log.Manager.Logger().SetExecutor(logger_term.New()) }

func initTracerFile()   { log.Manager.Tracer().SetExecutor(tracer_file.New()) }
func initTracerJaeger() { log.Manager.Tracer().SetExecutor(tracer_jaeger.New()) }
func initTracerTerm()   { log.Manager.Tracer().SetExecutor(tracer_term.New()) }
func initTracerZipkin() { log.Manager.Tracer().SetExecutor(tracer_zipkin.New()) }

func main() {
	log.Manager.Start(ctx)
	defer log.Manager.Stop()

	main1()
	main2()
	main3()

	time.Sleep(time.Second)
}

func main1() {
	log.Debug("debug message")
	log.Info("info message")
	log.Warn("warn message")
	log.Error("error message")
	log.Fatal("fatal message")
}

func main2() {
	log.Field{"key": "value"}.
		Fatal("fatal message")
}

func main3() {
	tr := tracers.NewTrace("trace")

	sp := tr.New("span")
	sp.Kv().Add("sk", "span value")
	defer sp.End()

	sp.Logger().Add("span-k", "span v").Info("span info 1 message")
	sp.Logger().Add("span-key", "span value").Fatal("span fatal 2 message")

	s2 := sp.Child("child")
	defer s2.End()

}

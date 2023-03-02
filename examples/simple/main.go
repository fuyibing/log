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

package main

import (
	"context"
	"github.com/fuyibing/log/v5"
	"github.com/fuyibing/log/v5/tracer"
	"time"
)

func main() {
	log.Logger.Start(context.Background())
	defer log.Logger.Stop()

	// main0()
	main1()
	// main2()
	// main3()

	time.Sleep(time.Second)
}

func main0() {}

func main1() {
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.Fatal("fatal")
}

func main2() {
	log.Field().
		Add("key", "value").
		Info("field info")
}

func main3() {
	tr := tracer.NewTrace("trace")

	sp := tr.Span("span")
	defer sp.End()
	sp.Logger().
		Add("sk", "span value").
		Info("span info")

}

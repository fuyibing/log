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
	"context"
	"github.com/fuyibing/log/v5"
	"time"
)

func init() {
}

func main() {

	ch := make(chan int)
	ctx := context.Background()

	go func() {
		_ = log.Manager.Start(ctx)
		ch <- 1
	}()

	go func() {
		time.Sleep(time.Second * 3)
		log.Manager.Stop()
	}()

	go func() {
		tr := log.Manager.NewTrace("trace")
		sp := tr.NewSpan("span")
		defer sp.End()

		sp.Logger().Info("trace id = %v", tr.GetTraceId().String())

		sp.GetAttr().
			Add("sk", "span value").
			Add("http.request.method", "POST").
			Add("http.request.url", "/ping")

		sp.Logger().Info("span info message")
		sp.Logger().Add("key", "value").Fatal("span fatal message")
	}()

	for {
		select {
		case <-ch:
			println("ended.")
			return
		}
	}
}

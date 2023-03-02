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
)

func main() {
	log.Logger.Start(context.Background())
	defer log.Logger.Stop()

	main0()
}

func main0() {
	trace := log.NewTrace("trace")

	span1 := trace.Span("span1")
	span1.Logger().Info("span 1: info level")
	span1.End()

	span21 := span1.Child("span21")
	span21.End()

	span22 := span1.Child("span22")
	span22.End()
}

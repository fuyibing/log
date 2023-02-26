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
	"github.com/fuyibing/log/v5/cores"
	"github.com/fuyibing/log/v5/exporters/logger_term"
	"testing"
	"time"
)

func TestField_Info(t *testing.T) {

	// Set logger exporter on: Terminal / Console.
	cores.Registry.RegisterLoggerExporter(
		logger_term.NewExporter(),
	)

	Info("info")
	Field{}.Info("field info: mapper mode")
	Field{}.
		Add("k1", "v1").
		Add("key2", "value2").Info("field info: Add mode")

	time.Sleep(time.Second)
}

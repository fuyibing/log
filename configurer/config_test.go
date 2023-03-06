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
// date: 2023-03-03

package configurer

import (
	"encoding/json"
	"github.com/fuyibing/log/v5/common"
	"testing"
)

func TestConfig_Setter(t *testing.T) {
	before, _ := json.Marshal(Config)
	t.Logf("config before: %s\n", before)

	Config.Setter().SetLoggerLevel(common.Debug)

	after, _ := json.Marshal(Config)
	t.Logf("config  after: %s\n", after)
}

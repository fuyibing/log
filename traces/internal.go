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

package traces

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// InternalError
// 打印内部错误.
func InternalError(text string, args ...interface{}) {
	var (
		file string
		line int
		list = []string{fmt.Sprintf(text, args...)}
		ln   = "\n"
		ok   bool
	)

	for i := 0; ; i++ {
		if _, file, line, ok = runtime.Caller(i); !ok {
			break
		}
		list = append(list, fmt.Sprintf("    #%d. %s:%d", i, strings.TrimSpace(file), line))
	}

	_, _ = fmt.Fprintf(os.Stdout, strings.Join(list, ln)+ln)
}

func InternalInfo(text string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, fmt.Sprintf(text, args...)+"\n")
}

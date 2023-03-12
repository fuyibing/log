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

package common

import (
	"fmt"
	"os"
)

var (
	ln = "\n"
)

// InternalInfo
// print INFO level to standard output device.
func InternalInfo(text string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, fmt.Sprintf(text, args...)+ln)
}

// InternalFatal
// print ERROR level to standard error device.
func InternalFatal(text string, args ...interface{}) {
	text = fmt.Sprintf(text, args...)

	// Range stacks and bound them on to message.
	for i, item := range Backstack().Items {
		text += fmt.Sprintf("\n%d. %s:%d call %s",
			i,
			item.File,
			item.Line,
			item.Call,
		)
	}

	_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", text))
}

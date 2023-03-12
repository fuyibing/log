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
// date: 2023-03-04

package common

import (
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
)

var (
	stackBegin           = "runtime/debug.Stack()"
	stackRegexFile       = regexp.MustCompile(`^([^:]+):(\d+)`)
	stackRegexFunc       = regexp.MustCompile(`([^/]+)$`)
	stackRegexFuncParams = regexp.MustCompile(`\([0{][^)]+\)`)
	stackRegexInternal   = regexp.MustCompile(`/fuyibing/log/v5[./]`)
)

type (
	Stack struct {
		Items []StackItem
	}

	StackItem struct {
		// Return true if code and line is current package.
		Internal bool

		// Return func name of stack position.
		Call string

		// File path name.
		File string

		// File line number.
		Line int
	}
)

// Backstack
// extract debug stack into list.
func Backstack() Stack {
	var (
		call       string
		lines      = strings.Split(string(debug.Stack()), "\n")
		stack      = Stack{Items: make([]StackItem, 0)}
		startIndex int
		startMax   = len(lines)
	)
	// Range stack to find start line number.
	for i, s := range lines {
		if s = strings.TrimSpace(s); s == "" {
			break
		}
		if s == stackBegin {
			startIndex = i + 2
			break
		}
	}
	// Range stack.
	for i := startIndex; i < startMax; i += 2 {
		if call = lines[i]; call == "" {
			break
		}
		// Init item struct.
		item := StackItem{Call: call, Internal: stackRegexInternal.MatchString(call)}
		// Remove func params.
		item.Call = stackRegexFuncParams.ReplaceAllString(item.Call, `()`)
		// Collect func name.
		if m := stackRegexFunc.FindStringSubmatch(item.Call); len(m) > 0 {
			item.Call = m[1]
		}
		// Parse file and line.
		if (i + 1) < startMax {
			if item.File = strings.TrimSpace(lines[i+1]); item.File != "" {
				if m := stackRegexFile.FindStringSubmatch(item.File); len(m) > 0 {
					item.File = m[1]
					if n, ne := strconv.ParseInt(m[2], 10, 32); ne == nil {
						item.Line = int(n)
					}
				}
			}
		}
		// Append to list.
		stack.Items = append(stack.Items, item)
	}
	return stack
}

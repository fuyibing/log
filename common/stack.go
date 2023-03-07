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
	// Stack
	// 栈结构.
	Stack struct {
		Items []StackItem
	}

	// StackItem
	// 栈元素.
	StackItem struct {
		// 是否内部.
		// true 表示元素是 fuyibing/log 包中的文件.
		Internal bool

		// 函数名称.
		// 例如: main.main()
		Call string

		// 文件路径.
		// 例如: /home/app/github.com/fuyibing/log/v5/field.go
		File string

		// 触发行号.
		Line int
	}
)

// Backstack
// 解析堆栈.
func Backstack() Stack {
	var (
		call       string
		lines      = strings.Split(string(debug.Stack()), "\n")
		stack      = Stack{Items: make([]StackItem, 0)}
		startIndex int
		startMax   = len(lines)
	)
	// 遍历堆栈.
	for i, s := range lines {
		if s = strings.TrimSpace(s); s == "" {
			break
		}
		if s == stackBegin {
			startIndex = i + 2
			break
		}
	}
	// 遍历堆栈.
	for i := startIndex; i < startMax; i += 2 {
		if call = lines[i]; call == "" {
			break
		}
		// 元素结构.
		item := StackItem{Call: call, Internal: stackRegexInternal.MatchString(call)}
		// 去除参数.
		item.Call = stackRegexFuncParams.ReplaceAllString(item.Call, `()`)
		// 提取名称.
		if m := stackRegexFunc.FindStringSubmatch(item.Call); len(m) > 0 {
			item.Call = m[1]
		}
		// 文件解析.
		// - 文件
		// - 行号
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
		// 堆栈列表.
		stack.Items = append(stack.Items, item)
	}
	return stack
}

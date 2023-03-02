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

package logger_file

import (
	"fmt"
	"github.com/fuyibing/log/v5/conf"
	"os"
	"sync"
	"time"
)

type (
	// writer
	// 写入文件.
	writer struct {
		caches map[string]bool
		sync.Mutex
	}
)

// 前置检查.
func (o *writer) check(t time.Time) (path, full string, err error) {
	o.Lock()
	defer o.Unlock()

	// 计算目录.
	path = fmt.Sprintf("%s/%s",
		conf.Config.GetFileLogger().GetPath(),
		t.Format(conf.Config.GetFileLogger().GetFolder()),
	)

	// 完整路径.
	full = fmt.Sprintf("%s/%s.%s",
		path,
		t.Format(conf.Config.GetFileLogger().GetName()),
		conf.Config.GetFileLogger().GetExt(),
	)

	// 重复检查.
	if _, ok := o.caches[path]; ok {
		return
	}

	// 创建目录.
	if err = os.MkdirAll(path, os.ModePerm); err == nil {
		o.caches[path] = true
	}
	return
}

// 文件构造.
func (o *writer) init() *writer {
	o.caches = make(map[string]bool)
	return o
}

// 写入文件.
func (o *writer) write(path string, text string) (err error) {
	var (
		fp *os.File
	)

	// 打开文件.
	if fp, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm); err != nil {
		return
	}

	// 关闭文件.
	defer func() { _ = fp.Close() }()

	// 写入日志.
	_, err = fp.WriteString(text)
	return
}

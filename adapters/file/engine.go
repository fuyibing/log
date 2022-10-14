// author: wsfuyibing <websearch@163.com>
// date: 2022-10-14

package file

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type (
	// Engine
	// 文件操作.
	Engine interface {
		Write(text string) (err error)
	}

	engine struct {
		err        error
		name, path string
		filepath   string
		mu         sync.Mutex
	}
)

// Write
// 写入内容.
func (o *engine) Write(text string) (err error) {
	var fp *os.File

	// 1. 准备写入.
	o.mu.Lock()
	defer func() {
		// 关闭文件.
		if fp != nil {
			_ = fp.Close()
		}

		// 捕获异常.
		if v := recover(); v != nil {
			err = fmt.Errorf("panic on file adapter: %v", v)
		}

		// 完成写入.
		o.mu.Unlock()
	}()

	// 2. 内部错误.
	if o.err != nil {
		return o.err
	}

	// 3. 打开文件.
	if fp, err = os.OpenFile(o.filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm); err != nil {
		return
	}

	// 4. 写入内容.
	_, err = fp.WriteString(fmt.Sprintf("%s\n", text))
	return
}

// 构造实例.
func (o *engine) init(t time.Time) *engine {
	o.mu = sync.Mutex{}
	o.name = fmt.Sprintf("%s.log", t.Format(Config.Name))
	o.path = fmt.Sprintf("%s/%s", Config.Path, t.Format(Config.Folder))
	o.filepath = fmt.Sprintf("%s/%s", o.path, o.name)
	o.err = os.MkdirAll(o.path, os.ModePerm)
	return o
}

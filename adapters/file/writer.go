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
	// Writer
	// 文件操作.
	Writer interface {
		Write(text string) (err error)
	}

	writer struct {
		err error
		mu  sync.Mutex

		name     string // 文件名, 如: 2020-10-01.log
		path     string // 路径名, 如: /var/logs/2020-10
		filepath string // 完整名, 如: /var/logs/2020-10/2020-10-01.log
	}
)

func (o *writer) Write(text string) (err error) {
	var (
		fp *os.File
	)

	o.mu.Lock()
	defer func() {
		if fp != nil {
			_ = fp.Close()
		}

		if v := recover(); v != nil {
			err = fmt.Errorf("%v", v)
		}

		o.mu.Unlock()
	}()

	if o.err != nil {
		return o.err
	}

	if fp, err = os.OpenFile(o.filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm); err != nil {
		return
	}

	_, err = fp.WriteString(fmt.Sprintf("%s\n", text))
	return
}

func (o *writer) init(t time.Time) *writer {
	o.mu = sync.Mutex{}
	o.name = fmt.Sprintf("%s.log", t.Format(Config.Name))
	o.path = fmt.Sprintf("%s/%s", Config.Path, t.Format(Config.Folder))
	o.filepath = fmt.Sprintf("%s/%s", o.path, o.name)
	o.err = os.MkdirAll(o.path, os.ModePerm)
	return o
}

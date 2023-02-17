// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package file

import (
	"fmt"
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/log/v8/formatters"
	"os"
	"strings"
	"sync"
	"time"
)

type (
	adapter struct {
		child     adapters.AdapterManager
		formatter formatters.Formatter
		ignorer   adapters.AdapterIgnore

		mu      *sync.RWMutex
		folders map[string]bool
	}
)

// /////////////////////////////////////////////////////////////
// Exported interface methods
// /////////////////////////////////////////////////////////////

func (o *adapter) Log(lines ...*base.Line)             { o.doLogs(lines...) }
func (o *adapter) SetChild(v adapters.AdapterManager)  { o.child = v }
func (o *adapter) SetFormatter(v formatters.Formatter) { o.formatter = v }
func (o *adapter) SetIgnore(v adapters.AdapterIgnore)  { o.ignorer = v }

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *adapter) doLogs(lines ...*base.Line) {
	var (
		err  error
		list       = make([]string, 0)
		u    int64 = 0
		ut         = time.Now()
	)

	defer func() {
		if err != nil {
			if o.child != nil {
				o.child.Log(lines...)
				return
			}

			if o.ignorer != nil {
				o.ignorer(err, lines...)
				return
			}
		}

		for _, line := range lines {
			line.Release()
		}
	}()

	for _, x := range lines {
		if u == 0 {
			ut = x.Time
			u = ut.Unix()
		} else if u > x.Time.Unix() {
			ut = x.Time
			u = ut.Unix()
		}

		list = append(list, o.formatter(x))
	}

	err = o.write(ut, strings.Join(list, "\n"))
}

func (o *adapter) init() *adapter {
	o.folders = make(map[string]bool)
	o.formatter = formatters.NewTextFormatter().Format
	o.mu = &sync.RWMutex{}
	return o
}

func (o *adapter) write(ut time.Time, text string) (err error) {
	var (
		fp     *os.File
		folder = fmt.Sprintf("%s/%s", conf.Config.File.Path, ut.Format(conf.Config.File.Folder))
		path   = fmt.Sprintf("%s/%s.%s", folder, ut.Format(conf.Config.File.Name), conf.Config.File.ExtName)
	)

	defer func() {
		if fp != nil {
			_ = fp.Close()
		}
	}()

	// Make directory.
	if err = func() error {
		o.mu.Lock()
		defer o.mu.Unlock()

		if _, ok := o.folders[folder]; ok {
			return nil
		}

		if me := os.MkdirAll(folder, os.ModePerm); me != nil {
			return me
		}

		o.folders[folder] = true
		return nil
	}(); err != nil {
		return
	}

	// Open file.
	if fp, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm); err != nil {
		return
	}

	// Write contents.
	_, err = fp.WriteString(fmt.Sprintf("%s\n", text))
	return
}

// /////////////////////////////////////////////////////////////
// Package init
// /////////////////////////////////////////////////////////////

var (
	Adapter *adapter
)

func init() {
	new(sync.Once).Do(func() {
		Adapter = (&adapter{}).init()
	})
}

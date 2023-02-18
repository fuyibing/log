// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package file

import (
	"fmt"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/log/v8/formatters"
	"os"
	"strings"
	"sync"
	"time"
)

type (
	Executor struct {
		formatter  formatters.Formatter
		mu         *sync.RWMutex
		separators map[string]time.Time
	}
)

func New() *Executor {
	return (&Executor{}).init()
}

// /////////////////////////////////////////////////////////////
// Exported methods
// /////////////////////////////////////////////////////////////

func (o *Executor) Logs(lines ...*base.Line) (err error) {
	var (
		fp             *os.File
		list           = make([]string, 0)
		path, fullPath string
		pathTime       time.Time
	)

	defer func() {
		if fp != nil {
			_ = fp.Close()
		}
	}()

	// Iterate lines.
	for i, line := range lines {
		// Call check process for first line.
		if i == 0 {
			pathTime = line.Time

			// Check path.
			if path, err = o.checkSeparator(pathTime); err != nil {
				return
			}
		}

		// Append to list.
		list = append(list, o.formatter.String(line))
	}

	// Generate full path name.
	fullPath = fmt.Sprintf("%s/%s.%s",
		path,
		pathTime.Format(conf.Config.GetFile().GetFileName()),
		conf.Config.GetFile().GetExtName(),
	)

	// Open file.
	if fp, err = os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm); err != nil {
		return
	}

	// Write log contents.
	_, err = fp.WriteString(fmt.Sprintf("%s\n", strings.Join(list, "\n")))
	return
}

func (o *Executor) SetFormatter(formatter formatters.Formatter) {
	o.formatter = formatter
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *Executor) checkSeparator(t time.Time) (path string, err error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	// Generation
	// separator directory.
	path = fmt.Sprintf("%s/%s",
		t.Format(conf.Config.GetFile().GetBasePath()),
		t.Format(conf.Config.GetFile().GetSeparatorPath()),
	)

	// Return
	// if separator directory exists.
	if _, ok := o.separators[path]; ok {
		return
	}

	// Return error
	// if make separator director failed.
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		return
	}

	// Update separators remarks.
	o.separators[path] = t
	return
}

func (o *Executor) init() *Executor {
	o.mu = &sync.RWMutex{}
	o.separators = make(map[string]time.Time)
	return o
}

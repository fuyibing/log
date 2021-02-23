// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package adapters

import (
	"sync"

	"github.com/fuyibing/log/interfaces"
)

// 文件配置.
type fileConfig struct {
	Path     string `yaml:"path"`
	UseMonth bool   `yaml:"use-month"`
}

type fileAdapter struct {
	Conf *fileConfig `yaml:"file"`
	ch   chan interfaces.LineInterface
	mu   *sync.RWMutex
}

func (o *fileAdapter) Run(line interfaces.LineInterface) {
	go func() {
		o.ch <- line
	}()
}

// Listen channel.
func (o *fileAdapter) listen() {
	go func() {
		defer o.listen()
		for {
			select {
			case line := <-o.ch:
				go o.send(line)
			}
		}
	}()
}

// Send log.
func (o *fileAdapter) send(line interfaces.LineInterface) {}

func NewFile() *fileAdapter {
	o := &fileAdapter{ch: make(chan interfaces.LineInterface), mu: new(sync.RWMutex)}
	o.listen()
	return o
}

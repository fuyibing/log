// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package base

import (
	"sync"
)

var (
	Pool PoolManager
)

type (
	PoolManager interface {
		AcquireLine() *Line
		ReleaseLine(x *Line)
	}

	pool struct {
		lines sync.Pool
	}
)

func (o *pool) AcquireLine() *Line {
	if v := o.lines.Get(); v != nil {
		return v.(*Line).before()
	}
	return (&Line{}).init().before()
}

func (o *pool) ReleaseLine(x *Line) {
	x.after()
	o.lines.Put(x)
}

func (o *pool) init() *pool {
	return o
}

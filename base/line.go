// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package base

import (
	"context"
	"github.com/fuyibing/log/v8/conf"
	"time"
)

type (
	// Line
	// log line definitions.
	Line struct {
		Ctx   context.Context
		Level conf.Level
		Text  string
		Time  time.Time
	}
)

func (o *Line) Release() { Pool.ReleaseLine(o) }

func (o *Line) after() *Line {
	return o
}

func (o *Line) before() *Line {
	o.Time = time.Now()
	return o
}

func (o *Line) init() *Line {
	return o
}

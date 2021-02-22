// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package adapters

import (
	"github.com/fuyibing/log/interfaces"
)

type redisAdapter struct {
}

func (o *redisAdapter) Run(lineInterface interfaces.LineInterface) {}

func NewRedis() *redisAdapter {
	return &redisAdapter{}
}
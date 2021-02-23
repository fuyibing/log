// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package tests

import (
	"testing"

	"github.com/fuyibing/log/v2/plugins"
)

func TestPluginXorm(t *testing.T) {
	x := plugins.NewXOrm()
	x.Infof("")
}

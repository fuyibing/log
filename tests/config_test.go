// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package tests

import (
	"testing"

	"github.com/fuyibing/log/v2"
)

func TestConfig(t *testing.T) {
	println("app addr:", log.Config.AppAddr())
	println("app name:", log.Config.AppName())
	println("debug on:", log.Config.DebugOn())
	println(" info on:", log.Config.InfoOn())
	println(" warn on:", log.Config.WarnOn())
	println("error on:", log.Config.ErrorOn())
}

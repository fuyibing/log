// author: wsfuyibing <websearch@163.com>
// date: 2022-10-26

package log

import (
	"github.com/fuyibing/log/v3/base"
	"testing"
)

func ExampleConfiguration_SetAdapter() {
	Config.SetAdapter(base.Term)
	Client.Start()

	Client.Debug("debug")
	Client.Infof("info")
	Client.Warnf("warn")
	Client.Errorf("error")
}

func TestConfiguration_SetAdapter(t *testing.T) {
	ExampleConfiguration_SetAdapter()
}

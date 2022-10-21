// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package tests

import (
	"testing"

	"github.com/fuyibing/log/v3"
)

func TestConfig(t *testing.T) {
	t.Logf("config: %+v", log.Config)
}

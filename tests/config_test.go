// author: wsfuyibing <websearch@163.com>
// date: 2021-02-11

package tests

import (
	"testing"

	"github.com/fuyibing/log"
)

func TestConfig(t *testing.T) {
	t.Logf("adapter       : {{%v}}", log.Config.Adapter)
	t.Logf("adapter name  : {{%v}}", log.Config.AdapterName)
	t.Logf("level         : {{%v}}", log.Config.Level)
	t.Logf("level name    : {{%v}}", log.Config.LevelName)
	t.Logf("time format   : {{%v}}", log.Config.TimeFormat)
	t.Logf("span id name  : {{%v}}", log.Config.NameSpanId)
	t.Logf("span ver name : {{%v}}", log.Config.NameSpanVersion)
	t.Logf("trace id name : {{%v}}", log.Config.NameTraceId)
	t.Logf("redis net     : {{%v}}", log.Config.Redis.Network)
	t.Logf("redis addr    : {{%v}}", log.Config.Redis.Addr)
	t.Logf("redis db      : {{%v}}", log.Config.Redis.Index)
}

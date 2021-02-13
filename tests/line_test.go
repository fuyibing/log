// author: wsfuyibing <websearch@163.com>
// date: 2021-02-11

package tests

import (
	"testing"

	"github.com/fuyibing/log"
)

func TestLine(t *testing.T) {
	ctx := log.NewContext()
	l1 := log.NewLine(ctx, log.LevelDebug, "[d=1.234]message")
	t.Logf("[%s][%s] %s", l1.Timeline(), l1.GetLevelText(), l1.String())
	l2 := log.NewLine(ctx, log.LevelInfo, "message 2")
	t.Logf("[%s][%s] %s", l2.Timeline(), l2.GetLevelText(), l2.String())
}

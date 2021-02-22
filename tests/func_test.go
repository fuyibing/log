// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package tests

import (
	"testing"

	"github.com/fuyibing/log"
)

func TestFunc(t *testing.T) {
	ctx := log.NewContext()
	log.Client.Debugfc(ctx, "debug fc")
	log.Client.Infofc(ctx, "info fc")
	log.Client.Warnfc(ctx, "warn fc")
	log.Client.Errorfc(ctx, "error fc")
}

// author: wsfuyibing <websearch@163.com>
// date: 2021-03-14

package tests

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/fuyibing/log/v2"
)

func TestConfig(t *testing.T) {

	log.Config.TimeFormat = "15:04:05"

	req, _ := http.NewRequest(http.MethodPost, "/index", nil)
	req.URL, _ = url.ParseRequestURI("/index")


	ctx := log.NewTracing().UseRequest(req)

	log.Client.Debugfc(ctx, "Debug format:")
	log.Client.Infofc(ctx, "Info  format:")
	log.Client.Warnfc(ctx, "Warn  format:")
	log.Client.Errorfc(ctx, "Error format:")

	t.Logf("debug: %v.", log.Config.DebugOn())
	t.Logf(" info: %v.", log.Config.InfoOn())
	t.Logf(" warn: %v.", log.Config.WarnOn())
	t.Logf("error: %v.", log.Config.ErrorOn())

}

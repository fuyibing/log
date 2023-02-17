// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package tests

import (
	"encoding/json"
	"github.com/fuyibing/log/v8"
	"github.com/fuyibing/log/v8/conf"
	"testing"
)

func TestConfig_Custom(t *testing.T) {
	log.Config.Set(
		conf.WithAdapter(conf.Kafka, conf.File, conf.Term),
		conf.WithLevel(conf.Error),
		conf.WithPrefix("prefix"),
		conf.WithService("myapp"),
		conf.WithTimeFormat("15:04:05.999"),
	)
	configPrint(t)
}

func TestConfig_Default(t *testing.T) {
	configPrint(t)
}

func configPrint(t *testing.T) {
	buf, _ := json.Marshal(log.Config)
	t.Logf("config default: %s", buf)
	t.Logf("config status [fatal=%v, error=%v, warn=%v, info=%v, debug=%v]",
		log.Config.FatalOn(),
		log.Config.ErrorOn(),
		log.Config.WarnOn(),
		log.Config.InfoOn(),
		log.Config.DebugOn(),
	)
}

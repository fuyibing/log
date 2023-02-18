// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package conf

import (
	"encoding/json"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := (&configuration{}).init()
	printConfig(t, cfg)
}

func TestDefault(t *testing.T) {
	cfg := (&configuration{}).init()
	cfg.initYaml()
	cfg.initDefaults()
	printConfig(t, cfg)
}

func TestConfiguration_Set(t *testing.T) {
	cfg := &configuration{}
	cfg.init()

	cfg.Set(
		SetLevel(Warn),
		SetTimeFormat("2006-01-02 15:04:05.999999"),

		SetFileBasePath("/var/logs/myapp"), // /var/logs/myapp
		SetFileSeparatorPath("06-01-02"),   // /var/logs/myapp/23-02-18
		SetFileFileName("06-01-02-15"),     // /var/logs/myapp/23-02-18/23-02-18-14
		SetFileExtName("log"),              // /var/logs/myapp/23-02-18/23-02-18-14.log

		SetTermColor(true),
	)

	printConfig(t, cfg)
}

func printConfig(t *testing.T, cfg Configuration) {
	buf, _ := json.MarshalIndent(cfg, "", "  ")
	t.Logf("configuration follows:\n%s", buf)

	t.Logf("config level: %s", cfg.GetLevel())
	t.Logf("config status [fatal=%v, error=%v, warn=%v, info=%v, debug=%v]",
		cfg.FatalOn(),
		cfg.ErrorOn(),
		cfg.WarnOn(),
		cfg.InfoOn(),
		cfg.DebugOn(),
	)

}

// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package log

import (
	"fmt"
	"github.com/fuyibing/log/v3/adapters/file"
	"github.com/fuyibing/log/v3/adapters/kafka"
	"github.com/fuyibing/log/v3/adapters/redis"
	"github.com/fuyibing/log/v3/adapters/term"
	"github.com/fuyibing/log/v3/base"
	"github.com/fuyibing/log/v3/trace"
	"gopkg.in/yaml.v3"
	"net"
	"os"
	"regexp"
)

// Config
// 基础配置.
var Config *configuration

// 基础配置结构体.
type configuration struct {
	// 适配器定义.
	Adapter []base.AdapterName   `yaml:"adapter"` // 适配器名称(log.yaml)
	Term    *term.Configuration  `yaml:"term"`    // 适配器配置(log.yaml.term)
	File    *file.Configuration  `yaml:"file"`    // 适配器配置(log.yaml.file)
	Redis   *redis.Configuration `yaml:"redis"`   // 适配器配置(log.yaml.redis)
	Kafka   *kafka.Configuration `yaml:"kafka"`   // 适配器配置(log.yaml.kafka)

	// 基础配置.
	Level      base.Level     `yaml:"-"`     // 级别类型(Exec)
	LevelName  base.LevelName `yaml:"level"` // 级别名称(log.yaml)
	TimeFormat string         `yaml:"time"`  // 时间格式(log.yaml)

	// 服务参数.
	Name    string `yaml:"name"`    // 服务名称(app.yaml)
	Port    int    `yaml:"port"`    // 服务端口(app.yaml)
	Version string `yaml:"version"` // 服务版本(app.yaml)

	TraceId      string `yaml:"trace-id"`
	ParentSpanId string `yaml:"parent-span-id"`
	SpanId       string `yaml:"span-id"`
	SpanVersion  string `yaml:"span-version"`

	// 级别状态
	debugOn, infoOn, warnOn, errorOn bool
}

// 构造实例.
func (o *configuration) init() *configuration {
	return o.scan().sync().status()
}

// 解析文件.
func (o *configuration) scan() *configuration {
	for _, fs := range [][]string{
		{"./tmp/log.yaml", "../tmp/log.yaml", "./config/log.yaml", "../config/log.yaml"},
		{"./tmp/app.yaml", "../tmp/app.yaml", "./config/app.yaml", "../config/app.yaml"},
	} {
		for _, f := range fs {
			if body, err := os.ReadFile(f); err == nil {
				if yaml.Unmarshal(body, o) == nil {
					break
				}
			}
		}
	}

	o.Level = o.LevelName.Level()
	return o
}

// 启动状态.
func (o *configuration) status() *configuration {
	switch o.Level {
	case base.Error:
		o.errorOn = true
	case base.Warn:
		o.errorOn = true
		o.warnOn = true
	case base.Info:
		o.errorOn = true
		o.warnOn = true
		o.infoOn = true
	case base.Debug:
		o.errorOn = true
		o.warnOn = true
		o.infoOn = true
		o.debugOn = true
	}
	return o
}

// 同步配置.
// 从 YAML 文件读取的参数覆盖 base 包下的参数.
func (o *configuration) sync() *configuration {
	// 适配器配置.

	if o.Term != nil {
		term.Config.Override(o.Term)
	}
	if o.File != nil {
		file.Config.Override(o.File)
	}
	if o.Redis != nil {
		redis.Config.Override(o.Redis)
	}
	if o.Kafka != nil {
		kafka.Config.Override(o.Kafka)
	}

	// 链路参数

	if o.TraceId != "" {
		trace.TracingTraceId = o.TraceId
	}
	if o.ParentSpanId != "" {
		trace.TracingParentSpanId = o.ParentSpanId
	}
	if o.SpanId != "" {
		trace.TracingSpanId = o.SpanId
	}
	if o.SpanVersion != "" {
		trace.TracingSpanVersion = o.SpanVersion
	}

	// 基础参数

	if o.TimeFormat != "" {
		base.LogTimeFormat = o.TimeFormat
	}

	// 服务参数.

	if o.Name != "" {
		base.LogName = o.Name
	}
	if o.Port > 0 {
		base.LogPort = o.Port
	}
	if o.Version != "" {
		base.LogVersion = o.Version
	}
	base.LogUserAgent = fmt.Sprintf("%s/%s", base.LogName, base.LogVersion)

	// 部署节点.
	// 解析部署机器的IP地址.
	if nis, e1 := net.Interfaces(); e1 == nil {
		for _, ni := range nis {
			if list, e2 := ni.Addrs(); e2 == nil {
				for _, addr := range list {
					if m := regexp.MustCompile(`^(\d+\.\d+\.\d+\.\d+)`).FindStringSubmatch(addr.String()); len(m) > 0 && m[0] != "127.0.0.1" {
						base.LogHost = m[1]
						break
					}
				}
			}
		}
	}

	return o
}

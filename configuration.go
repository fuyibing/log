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
var Config *Configuration

// Configuration
// 基础配置结构体.
type Configuration struct {
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

// /////////////////////////////////////////////////////////////
// 状态开关
// /////////////////////////////////////////////////////////////

func (o *Configuration) DebugOn() bool {
	return o.debugOn
}

func (o *Configuration) ErrorOn() bool {
	return o.errorOn
}

func (o *Configuration) InfoOn() bool {
	return o.infoOn
}

func (o *Configuration) WarnOn() bool {
	return o.warnOn
}

// /////////////////////////////////////////////////////////////
// 调整配置
// /////////////////////////////////////////////////////////////

func (o *Configuration) SetAdapter(adapter ...base.Adapter) *Configuration {
	o.useAdapter(adapter...)
	return o
}

func (o *Configuration) SetLevel(level base.Level) *Configuration {
	o.useLevel(level)
	return o
}

// 构造实例.
func (o *Configuration) init() *Configuration {
	// 1. 解析配置.
	o.scan()

	// 2. 同步配置.
	o.syncAdapterConfig()
	o.syncBase()
	o.syncOpenTracing()

	// 3. 启用级别.
	if o.Level == base.Off {
		o.Level = base.Info
	}
	o.useLevel(o.Level)

	if o.Adapter == nil || len(o.Adapter) == 0 {
		o.useAdapter(base.Term)
	} else {
		o.useAdapterName(o.Adapter...)
	}

	return o
}

// 解析文件.
func (o *Configuration) scan() {
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
}

func (o *Configuration) syncAdapterConfig() {
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
}

func (o *Configuration) syncBase() *Configuration {
	// 1. 解析配置.

	if o.Name != "" {
		base.LogName = o.Name
	}
	if o.Port > 0 {
		base.LogPort = o.Port
	}
	if o.TimeFormat != "" {
		base.LogTimeFormat = o.TimeFormat
	}
	if o.Version != "" {
		base.LogVersion = o.Version
	}

	// 2. 动态计算.

	base.LogUserAgent = fmt.Sprintf("%s/%s", base.LogName, base.LogVersion)

	// 3. 动态计算.
	//    解析部署机器的IP地址.
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

func (o *Configuration) syncOpenTracing() {
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
}

func (o *Configuration) useAdapter(adapters ...base.Adapter) {
	var a base.AdapterEngine

	// 1. 遍历适配器.
	for _, adapter := range adapters {
		switch adapter {
		case base.Kafka:
			a = kafka.New().Parent(a)
		case base.Redis:
			a = redis.New().Parent(a)
		case base.File:
			a = file.New().Parent(a)
		case base.Term:
			a = term.New().Parent(a)
		}
	}

	// 2. 绑定适配器.
	Client.adapter = a
	Client.adapterOn = a != nil
}

func (o *Configuration) useAdapterName(names ...base.AdapterName) {
	adapters := make([]base.Adapter, 0)
	for _, name := range names {
		adapters = append(adapters, name.Adapter())
	}
	o.useAdapter(adapters...)
}

func (o *Configuration) useLevel(level base.Level) {
	o.Level = level

	o.debugOn = level >= base.Debug
	o.infoOn = level >= base.Info
	o.warnOn = level >= base.Warn
	o.errorOn = level >= base.Error
}

// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package log

import (
	"fmt"
	"net"
	"os"
	"regexp"

	"github.com/fuyibing/log/v3/adapters/file"
	"github.com/fuyibing/log/v3/adapters/kafka"
	"github.com/fuyibing/log/v3/adapters/redis"
	"github.com/fuyibing/log/v3/adapters/term"
	"github.com/fuyibing/log/v3/base"
	"gopkg.in/yaml.v3"
)

var Config *Configuration

type (
	Configuration struct {
		// 1. 服务定义
		//    从 config/app.yaml 中解析.

		Name    string `yaml:"name"`    // 服务名
		Port    int    `yaml:"port"`    // 端口号
		Version string `yaml:"version"` // 版本号

		// 2. 基础配置
		//    从 config/log.yaml 中解析.

		Adapter     base.Adapter     `yaml:"-"`       // 适配器类型
		AdapterName base.AdapterName `yaml:"adapter"` // 适配器名称
		Level       base.Level       `yaml:"-"`       // 级别类型
		LevelName   base.LevelName   `yaml:"level"`   // 级别名称
		TimeFormat  string           `yaml:"time"`    // 时间格式.

		// 3. 适配配置.
		//    从 config/log.yaml 中解析.

		Term  *term.Configuration  `yaml:"term"`  // 终端配置
		File  *file.Configuration  `yaml:"file"`  // 文件配置
		Redis *redis.Configuration `yaml:"redis"` // Redis 配置
		Kafka *kafka.Configuration `yaml:"kafka"` // Kafka 配置

		// n. 级别状态.
		debugOn, infoOn, warnOn, errorOn bool
	}
)

// /////////////////////////////////////////////////////////////
// 状态检测.
// /////////////////////////////////////////////////////////////

func (o *Configuration) DebugOn() bool { return o.debugOn }
func (o *Configuration) InfoOn() bool  { return o.infoOn }
func (o *Configuration) WarnOn() bool  { return o.warnOn }
func (o *Configuration) ErrorOn() bool { return o.errorOn }

// /////////////////////////////////////////////////////////////
// 配置解析.
// /////////////////////////////////////////////////////////////

// 赋默认值.
func (o *Configuration) defaults() *Configuration {
	if o.Name != "" {
		base.LogName = o.Name
	}
	if o.Port > 0 {
		base.LogPort = o.Port
	}
	if o.Version != "" {
		base.LogVersion = o.Version
	}
	if o.TimeFormat != "" {
		base.LogTimeFormat = o.TimeFormat
	}

	base.LogUserAgent = fmt.Sprintf("%s/%s", base.LogName, base.LogVersion)
	return o
}

// 构造实例.
func (o *Configuration) init() *Configuration {
	return o.runtime().scan().defaults().status()
}

// 运行参数.
func (o *Configuration) runtime() *Configuration {
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

// 解析文件.
func (o *Configuration) scan() *Configuration {
	for _, f := range []string{"./tmp/log.yaml", "../tmp/log.yaml", "./config/log.yaml", "../config/log.yaml"} {
		if body, err := os.ReadFile(f); err == nil {
			if yaml.Unmarshal(body, o) == nil {
				break
			}
		}
	}

	for _, f := range []string{"./tmp/app.yaml", "../tmp/app.yaml", "./config/app.yaml", "../config/app.yaml"} {
		if body, err := os.ReadFile(f); err == nil {
			if yaml.Unmarshal(body, o) == nil {
				break
			}
		}
	}

	return o
}

// 日志状态.
func (o *Configuration) status() *Configuration {
	// 1. 日志级级.
	o.Level = o.LevelName.Level()
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

	// 2. 适配器配置.
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

	// 3. 适配器绑定.
	//    若日志级别为关闭, 则不启动适配器.
	o.Adapter = o.AdapterName.Adapter()
	switch o.Adapter {
	case base.Term:
		Client.SetAdapter(term.New())
	case base.File:
		Client.SetAdapter(file.New().Parent(term.New()))
	case base.Redis:
		Client.SetAdapter(redis.New().Parent(file.New().Parent(term.New())))
	case base.Kafka:
		Client.SetAdapter(kafka.New().Parent(file.New().Parent(term.New())))
	}
	return o
}

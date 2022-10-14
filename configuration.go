// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package log

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/fuyibing/log/v3/adapters"
	"github.com/fuyibing/log/v3/adapters/file"
	"github.com/fuyibing/log/v3/adapters/kafka"
	"github.com/fuyibing/log/v3/adapters/redis"
	"github.com/fuyibing/log/v3/adapters/term"
	"gopkg.in/yaml.v3"
)

var Config *config

// 全局配置.
type config struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Port    int    `yaml:"port"`

	Adapter    adapters.Adapter  `yaml:"adapter"`
	Level      adapters.Level    `yaml:"-"`
	LevelKey   adapters.LevelKey `yaml:"level"`
	TimeFormat string            `yaml:"time"`

	Term  *term.Configuration  `yaml:"term"`
	Redis *redis.Configuration `yaml:"redis"`
	File  *file.Configuration  `yaml:"file"`
	Kafka *kafka.Configuration `yaml:"kafka"`

	debugOn, infoOn, warnOn, errorOn bool
}

// /////////////////////////////////////////////////////////////
// 状态级别.
// /////////////////////////////////////////////////////////////

func (o *config) DebugOn() bool { return o.debugOn }
func (o *config) InfoOn() bool  { return o.infoOn }
func (o *config) WarnOn() bool  { return o.warnOn }
func (o *config) ErrorOn() bool { return o.errorOn }

// 赋默认值.
func (o *config) defaults() *config {
	// 1. 适配器变量.
	if o.TimeFormat != "" {
		adapters.TimeFormat = o.TimeFormat
	}

	// 2. 日志级别.
	if o.LevelKey != "" {
		if level, ok := adapters.LevelKeys[o.LevelKey]; ok {
			o.Level = level
		}
	}

	// 3. 级别开关.
	switch o.Level {
	case adapters.Debug:
		o.debugOn = true
		o.infoOn = true
		o.warnOn = true
		o.errorOn = true
	case adapters.Info:
		o.debugOn = false
		o.infoOn = true
		o.warnOn = true
		o.errorOn = true
	case adapters.Warn:
		o.debugOn = false
		o.infoOn = false
		o.warnOn = true
		o.errorOn = true
	case adapters.Error:
		o.debugOn = false
		o.infoOn = false
		o.warnOn = false
		o.errorOn = true
	default:
		o.Level = adapters.Off
		o.debugOn = false
		o.infoOn = false
		o.warnOn = false
		o.errorOn = false
	}

	// 4. 创建配器.
	var (
		at  = term.New()
		atc = func() {
			if o.Term != nil {
				term.Config.Defaults(o.Term)
			}
			at.Start()
		}

		af  = file.New().Interrupt(at.Run)
		afc = func() {
			if o.File != nil {
				file.Config.Defaults(o.File)
			}
			af.Start()
		}

		ar  = redis.New().Interrupt(af.Run)
		arc = func() {
			if o.Redis != nil {
				redis.Config.Defaults(o.Redis)
			}
			ar.Start()
		}

		ak  = kafka.New().Interrupt(af.Run)
		akc = func() {
			if o.Kafka != nil {
				kafka.Config.Defaults(o.Kafka)
			}
			ar.Start()
		}
	)

	if o.Adapter == adapters.AdapterTerm {
		Client.setHandler(at.Run, atc)
	} else if o.Adapter == adapters.AdapterFile {
		Client.setHandler(af.Run, atc, afc)
	} else if o.Adapter == adapters.AdapterRedis {
		Client.setHandler(ar.Run, atc, afc, arc)
	} else if o.Adapter == adapters.AdapterKafka {
		Client.setHandler(ak.Run, atc, afc, akc)
	} else {
		panic(fmt.Sprintf("unknown %s logger adapter", o.Adapter))
	}
	return o
}

// 构造实例.
func (o *config) init() *config {
	return o.scan().defaults()
}

// 扫描配置.
// 读取配置文件并赋值给实例.
func (o *config) scan() *config {
	for _, f := range []string{"./tmp/log.yaml", "../tmp/log.yaml", "./config/log.yaml", "../config/log.yaml"} {
		if body, err := os.ReadFile(f); err == nil {
			if yaml.Unmarshal(body, o) == nil {
				break
			}
		}
	}
	return o.scanApp().scanHost()
}

func (o *config) scanApp() *config {
	for _, f := range []string{"./tmp/app.yaml", "../tmp/app.yaml", "./config/app.yaml", "../config/app.yaml"} {
		if body, err := os.ReadFile(f); err == nil {
			if yaml.Unmarshal(body, o) == nil {
				break
			}
		}
	}

	if o.Name != "" || o.Version != "" {
		if o.Name != "" {
			adapters.Name = o.Name
		}
		if o.Version != "" {
			adapters.Version = o.Version
		}
		adapters.Software = fmt.Sprintf("%s/%s", adapters.Name, adapters.Version)
	}

	if o.Port > 0 {
		adapters.Port = o.Port
	}

	return o
}

func (o *config) scanHost() *config {
	adapters.Pid = os.Getpid()

	if nis, e1 := net.Interfaces(); e1 == nil {
		for _, ni := range nis {
			if list, e2 := ni.Addrs(); e2 == nil {
				for _, addr := range list {
					if m := regexp.MustCompile(`^(\d+\.\d+\.\d+\.\d+)`).FindStringSubmatch(addr.String()); len(m) > 0 && m[0] != "127.0.0.1" {
						adapters.Host = m[1]
						break
					}
				}
			}
		}
	}

	adapters.NodeId = strings.ReplaceAll(fmt.Sprintf("%s_%d_%d", adapters.Host, rand.Intn(9999), rand.Intn(9999)), ".", "_")
	return o
}

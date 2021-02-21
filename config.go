// author: wsfuyibing <websearch@163.com>
// date: 2021-02-10

package log

import (
	"io/ioutil"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	regexpAddrPort     = regexp.MustCompile(`:(\d+)$`)
	regexpEthernetIpV4 = regexp.MustCompile(`^(\d+\.\d+\.\d+\.\d+)`)
)

// Configuration struct.
type configuration struct {
	Adapter         Adapter             `yaml:"-"`
	AdapterName     string              `yaml:"adapter"`
	Level           Level               `yaml:"-"`
	LevelName       string              `yaml:"level"`
	TimeFormat      string              `yaml:"time"`
	NameSpanId      string              `yaml:"span-id"`
	NameSpanVersion string              `yaml:"span-version"`
	NameTraceId     string              `yaml:"trace-id"`
	Redis           *configurationRedis `yaml:"redis"`
	Pid             int                 `yaml:"-"`
	ServerAddr      string              `yaml:"-"`
	ServerName      string              `yaml:"-"`
	ServerPort      int32               `yaml:"-"`
}

// Configuration for redis.
type configurationRedis struct {
	Addr            string `yaml:"addr"`
	Index           int    `yaml:"index"`
	Network         string `yaml:"network"`
	Password        string `yaml:"password"`
	MaxActive       int    `yaml:"max-active"`
	MaxIdle         int    `yaml:"max-idle"`
	Wait            bool   `yaml:"wait"`
	IdleTimeout     int    `yaml:"idle-timeout"`
	MaxConnLifetime int    `yaml:"max-conn-lifetime"`
}

// Load from YAML file.
func (o *configuration) LoadYaml(path string) error {
	var body []byte
	var err error
	// 1. read YAML file.
	body, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	// 2. parse YAML file.
	if err = yaml.Unmarshal(body, o); err != nil {
		return err
	}
	// 3. default: adapter.
	adapterDefault := true
	if o.AdapterName != "" {
		o.AdapterName = strings.ToUpper(o.AdapterName)
		for adapter, name := range AdapterText {
			if name == o.AdapterName {
				adapterDefault = false
				o.Adapter = adapter
				break
			}
		}
		if adapterDefault {
			panic("unknown log adapter")
		}
	}
	if adapterDefault {
		o.Adapter = DefaultAdapter
		o.AdapterName = AdapterText[o.Adapter]
	}
	// 4. default: level.
	levelDefault := true
	if o.LevelName != "" {
		o.LevelName = strings.ToUpper(o.LevelName)
		for level, name := range LevelText {
			if name == o.LevelName {
				levelDefault = false
				o.Level = level
				break
			}
		}
		if levelDefault {
			panic("unknown log level")
		}
	}
	if levelDefault {
		o.Level = DefaultLevel
		o.LevelName = LevelText[o.Level]
	}
	// 5. default: time format.
	if o.TimeFormat == "" {
		o.TimeFormat = DefaultTimeFormat
	}
	// 6. default: for open tracing.
	if o.NameSpanId == "" {
		o.NameSpanId = DefaultSpanId
	}
	if o.NameSpanVersion == "" {
		o.NameSpanVersion = DefaultSpanVersion
	}
	if o.NameTraceId == "" {
		o.NameTraceId = DefaultTraceId
	}
	return nil
}

// Initialize configuration.
func (o *configuration) initialize() {
	// 1. Process id.
	o.Pid = os.Getpid()
	// 2. Server ethernet IPv4.
	o.initializeAddr()
	// 3. Server name & port.
	o.initializeInfo()
	// 4. Log fields.
	for _, path := range []string{"./tmp/log.yaml", "./config/log.yaml", "../config/log.yaml"} {
		if err := o.LoadYaml(path); err == nil {
			break
		}
	}
}

// Initialize ethernet address.
// format: IPv4.
func (o *configuration) initializeAddr() {
	o.ServerAddr = "0.0.0.0"
	// Loop net interface.
	if nis, e1 := net.Interfaces(); e1 == nil {
		for _, ni := range nis {
			// Filtered by name.
			if ni.Name != "en0" && ni.Name != "eth0" {
				continue
			}
			// Point
			if addrs, e2 := ni.Addrs(); e2 == nil {
				for _, addr := range addrs {
					if m := regexpEthernetIpV4.FindStringSubmatch(addr.String()); len(m) > 0 {
						o.ServerAddr = m[1]
						break
					}
				}
			}
		}
	}
}

// Initialize server name and port.
func (o *configuration) initializeInfo() {
	// Parse sever name & port from yaml.
	var x = &struct {
		Addr string `yaml:"addr"`
		Name string `yaml:"name"`
	}{}
	// Load from file.
	for _, file := range []string{"./tmp/app.yaml", "./config/app.yaml", "../config/app.yaml"} {
		bs, e1 := ioutil.ReadFile(file)
		if e1 != nil {
			continue
		}
		if e2 := yaml.Unmarshal(bs, x); e2 != nil {
			continue
		}
	}
	// Port.
	o.ServerPort = 0
	if x.Addr != "" {
		if m := regexpAddrPort.FindStringSubmatch(x.Addr); len(m) == 2 {
			if n, e3 := strconv.ParseInt(m[1], 0, 32); e3 == nil {
				o.ServerPort = int32(n)
			}
		}
	}
	// Name.
	o.ServerName = "Unknown"
	if x.Name != "" {
		o.ServerName = x.Name
	}
}

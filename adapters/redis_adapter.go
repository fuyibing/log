// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package adapters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
	"gopkg.in/yaml.v3"

	"github.com/fuyibing/log/v2/interfaces"
)

type redisConfig struct {
	Address      string `yaml:"addr"`
	Database     int    `yaml:"database"`
	IdleTimeout  int    `yaml:"idle-timeout"`
	KeepAlive    int    `yaml:"keep-alive"`
	MaxActive    int    `yaml:"max-active"`
	MaxIdle      int    `yaml:"max-idle"`
	MaxLifetime  int    `yaml:"max-lifetime"`
	Network      string `yaml:"network"`
	Password     string `yaml:"password"`
	ReadTimeout  int    `yaml:"read-timeout"`
	Timeout      int    `yaml:"timeout"`
	Wait         bool   `yaml:"wait"`
	WriteTimeout int    `yaml:"write-timeout"`
	KeyLifetime  int    `yaml:"key-lifetime"`
	KeyPrefix    string `yaml:"key-prefix"`
	KeyList      string `yaml:"key-list"`
}

type redisAdapter struct {
	Conf    *redisConfig `yaml:"redis"`
	ch      chan interfaces.LineInterface
	pool    *redis.Pool
	handler interfaces.Handler
}

func (o *redisAdapter) Run(line interfaces.LineInterface) {
	go func() {
		o.ch <- line
	}()
}

func (o *redisAdapter) body(line interfaces.LineInterface) string {
	// Init
	data := make(map[string]interface{})
	// Basic.
	data["content"] = line.Content()
	data["duration"] = line.Duration()
	data["level"] = line.Level()
	data["time"] = line.Timeline()
	// Tracing.
	data["action"] = ""
	if line.Tracing() {
		data["parentSpanId"] = line.ParentSpanId()
		data["requestId"] = line.TraceId()
		data["requestMethod"], data["requestUrl"] = line.RequestInfo()
		data["spanId"] = line.SpanId()
		data["traceId"] = line.TraceId()
		data["version"] = line.SpanVersion()
	} else {
		data["parentSpanId"] = ""
		data["requestId"] = ""
		data["requestMethod"] = ""
		data["requestUrl"] = ""
		data["spanId"] = ""
		data["traceId"] = ""
		data["version"] = ""
	}
	// Server.
	data["module"] = line.ServiceName()
	data["pid"] = line.Pid()
	data["serverAddr"] = line.ServiceAddr()
	data["taskId"] = 0
	data["taskName"] = ""
	// JSON string.
	if body, err := json.Marshal(data); err == nil {
		return string(body)
	}
	return ""
}

func (o *redisAdapter) listen() {
	go func() {
		defer o.listen()
		for {
			select {
			case line := <-o.ch:
				go o.send(line)
			}
		}
	}()
}

func (o *redisAdapter) send(line interfaces.LineInterface) {
	// Catch panic.
	defer func() {
		if r := recover(); r != nil {
			o.handler(line)
		}
	}()
	// Pool: get & release.
	p := o.pool.Get()
	defer func() {
		_ = p.Close()
	}()
	// Send command.
	data := o.body(line)
	keyList := fmt.Sprintf("%s:%s", o.Conf.KeyPrefix, o.Conf.KeyList)
	keyItem := fmt.Sprintf("%s:t%10d%09d%12d", o.Conf.KeyPrefix, line.Time().Unix(), line.Time().Nanosecond(), rand.Int63n(999999999999))
	if err := p.Send("SET", keyItem, data, "EX", o.Conf.KeyLifetime); err == nil {
		_ = p.Send("RPUSH", keyList, keyItem)
	}
}

// 创建适配器.
func NewRedis() *redisAdapter {
	o := &redisAdapter{ch: make(chan interfaces.LineInterface), handler: NewTerm().Run}
	// Parse configuration.
	// 1. base config.
	for _, file := range []string{"./tmp/log.yaml", "../tmp/log.yaml", "./config/log.yaml", "../config/log.yaml"} {
		body, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}
		if yaml.Unmarshal(body, o) != nil {
			continue
		}
		break
	}
	// 2. key settings.
	if o.Conf.KeyLifetime == 0 {
		o.Conf.KeyLifetime = 7200
	}
	if o.Conf.KeyPrefix == "" {
		o.Conf.KeyPrefix = "logger"
	}
	if o.Conf.KeyList == "" {
		o.Conf.KeyList = "list"
	}
	// 3. Redis pool.
	o.pool = &redis.Pool{MaxActive: o.Conf.MaxActive, MaxIdle: o.Conf.MaxIdle, Wait: o.Conf.Wait}
	// 3.1 lifetime
	if o.Conf.MaxLifetime > 0 {
		o.pool.MaxConnLifetime = time.Duration(o.Conf.MaxLifetime) * time.Second
	}
	// 3.2 timeout: idle
	if o.Conf.IdleTimeout > 0 {
		o.pool.IdleTimeout = time.Duration(o.Conf.IdleTimeout) * time.Second
	}
	// 3.3 Connect
	o.pool.Dial = func() (redis.Conn, error) {
		// options: default.
		opts := make([]redis.DialOption, 0)
		opts = append(opts, redis.DialPassword(o.Conf.Password), redis.DialDatabase(o.Conf.Database))
		// options: timeouts.
		//          connect
		//          read
		//          write
		if o.Conf.Timeout > 0 {
			opts = append(opts, redis.DialConnectTimeout(time.Duration(o.Conf.Timeout)*time.Second))
		}
		if o.Conf.ReadTimeout > 0 {
			opts = append(opts, redis.DialReadTimeout(time.Duration(o.Conf.ReadTimeout)*time.Second))
		}
		if o.Conf.WriteTimeout > 0 {
			opts = append(opts, redis.DialWriteTimeout(time.Duration(o.Conf.WriteTimeout)*time.Second))
		}
		// options: keep alive
		if o.Conf.KeepAlive > 0 {
			opts = append(opts, redis.DialKeepAlive(time.Duration(o.Conf.KeepAlive)*time.Second))
		}
		// create connection
		return redis.Dial(o.Conf.Network, o.Conf.Address, opts...)
	}
	o.listen()
	return o
}

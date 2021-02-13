// author: wsfuyibing <websearch@163.com>
// date: 2021-02-10

package log

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Redis struct.
type (
	adapterRedis struct {
		ch   chan LineInterface
		pool *redis.Pool
	}
)

// New redis adapter instance.
func NewAdapterRedis() *adapterRedis {
	o := &adapterRedis{ch: make(chan LineInterface)}
	o.listen()
	o.pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				Config.Redis.Network,
				Config.Redis.Addr,
				redis.DialPassword(Config.Redis.Password),
				redis.DialDatabase(Config.Redis.Index),
			)
		},
		Wait:            Config.Redis.Wait,
		MaxActive:       Config.Redis.MaxActive,
		MaxIdle:         Config.Redis.MaxIdle,
		IdleTimeout:     time.Duration(Config.Redis.IdleTimeout) * time.Second,
		MaxConnLifetime: time.Duration(Config.Redis.MaxConnLifetime) * time.Second,
	}
	return o
}

// Send line to channel when called.
func (o *adapterRedis) Callback(line LineInterface) {
	go func() {
		o.ch <- line
	}()
}

// Format log body.
func (o *adapterRedis) format(line LineInterface) (string, error) {
	// Init
	data := make(map[string]interface{})
	// Basic.
	data["content"] = line.String()
	data["duration"] = line.Duration()
	data["level"] = line.GetLevelText()
	data["time"] = line.Timeline()
	// Tracing.
	data["action"] = ""
	if t := line.Tracing(); t != nil {
		data["parentSpanId"] = t.ParentSpanId()
		data["requestId"] = t.TraceId()
		data["requestMethod"] = t.RequestMethod()
		data["requestUrl"] = t.RequestUrl()
		data["spanId"] = t.SpanId()
		data["traceId"] = t.TraceId()
		data["version"] = t.Version(line.TracingOffset())
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
	data["module"] = Config.ServerName
	data["pid"] = Config.Pid
	data["serverAddr"] = fmt.Sprintf("%s:%d", Config.ServerAddr, Config.ServerPort)
	data["taskId"] = 0
	data["taskName"] = ""
	// JSON
	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Receive channel data.
func (o *adapterRedis) listen() {
	go func() {
		defer o.listen()
		for line := range o.ch {
			go o.send(line)
		}
	}()
}

// Send data to redis.
func (o *adapterRedis) send(line LineInterface) {
	// parse body and return if error.
	b, err := o.format(line)
	if err != nil {
		return
	}
	// key initialize.
	t := line.Time()
	k := fmt.Sprintf("logger:a%d%d%d", t.Unix(), t.UnixNano(), rand.Int63n(999999999999))
	l := "logger:list"
	// Redis pool.
	cli := o.pool.Get()
	defer func() {
		_ = cli.Close()
	}()
	// Send to redis.
	if err = cli.Send("SET", k, b, "NX", "EX", 3600); err == nil {
		_, _ = cli.Do("LPUSH", l, k)
	}
}

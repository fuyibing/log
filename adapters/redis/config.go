// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package redis

var Config *Configuration

// Configuration
// Redis适配器配置.
type Configuration struct {
	// 并发限流.
	Concurrency int32 `yaml:"concurrency"` // 最大并发.

	// 批量上报.
	// 每批最多上报多少条日志.
	Limit int `yaml:"limit"` // 每批提交 100 条.

	// 定时检查.
	// 每隔多长时间(单位: 毫秒)检查是否有未提交的日志, 若有未提交日志
	// 则立即上报.
	Ticker int `yaml:"ticker"`

	// Redis键名.

	KeyPrefix   string `yaml:"key-prefix"`
	KeyLifetime int    `yaml:"key-lifetime"`
	KeyList     string `yaml:"key-list"`

	// Redis连接.

	Network      string `yaml:"network"`
	Address      string `yaml:"addr"`
	Password     string `yaml:"password"`
	Database     int    `yaml:"database"`
	MaxActive    int    `yaml:"max-active"`    // 最大活跃连接
	MaxIdle      int    `yaml:"max-idle"`      // 最大空闲连接
	Timeout      int    `yaml:"timeout"`       // 连接超时时长(等待连接结果返回)
	ReadTimeout  int    `yaml:"read-timeout"`  // 读取超时时长(等待读取结果返回)
	WriteTimeout int    `yaml:"write-timeout"` // 写入超时时长(等待写入结果返回)

	Wait *bool `yaml:"wait"`
}

// Defaults
// 覆盖配置.
func (o *Configuration) Defaults(x *Configuration) {
	// 并发限流.

	if o.Concurrency == 0 {
		o.Concurrency = 5
	}
	if o.Limit == 0 {
		o.Limit = 100
	}
	if o.Ticker == 0 {
		o.Ticker = 500
	}

	// Redis键名.

	if x.KeyPrefix != "" {
		o.KeyPrefix = x.KeyPrefix
	}
	if x.KeyList != "" {
		o.KeyList = x.KeyList
	}
	if x.KeyLifetime > 0 {
		o.KeyLifetime = x.KeyLifetime
	}

	// Redis连接.

	if x.Network != "" {
		o.Network = x.Network
	}
	if x.Address != "" {
		o.Address = x.Address
	}
	if x.Password != "" {
		o.Password = x.Password
	}
	if x.Database > 0 {
		o.Database = x.Database
	}
	if x.MaxActive > 0 {
		o.MaxActive = x.MaxActive
	}
	if x.MaxIdle > 0 {
		o.MaxIdle = x.MaxIdle
	}
	if x.Timeout > 0 {
		o.Timeout = x.Timeout
	}
	if x.ReadTimeout > 0 {
		o.ReadTimeout = x.ReadTimeout
	}
	if x.WriteTimeout > 0 {
		o.WriteTimeout = x.WriteTimeout
	}
	if x.Wait != nil {
		o.Wait = x.Wait
	}
}

func (o *Configuration) init() *Configuration {
	o.KeyPrefix = "logger"
	o.KeyLifetime = 7200
	o.KeyList = "list"

	o.Network = "tcp"
	o.Address = "127.0.0.1:6379"
	o.MaxActive = 5
	o.MaxIdle = 2

	o.Timeout = 5
	o.ReadTimeout = 3
	o.WriteTimeout = 10

	yes := true
	o.Wait = &yes

	return o
}

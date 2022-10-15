// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package redis

var Config *Configuration

const (
	defaultConcurrency    int32 = 10
	defaultBatchFrequency       = 200
	defaultBatchLimit           = 100

	defaultKeyPrefix   = "logger"
	defaultKeyLifetime = 3600
	defaultKeyList     = "list"

	defaultNetwork      = "tcp"
	defaultAddress      = "127.0.0.1:6379"
	defaultMaxActive    = 5
	defaultMaxIdle      = 2
	defaultTimeout      = 5
	defaultReadTimeout  = 2
	defaultWriteTimeout = 10
)

var (
	defaultPoolWait = true
	defaultPoolNode = ""
)

// Configuration
// 基础配置.
type Configuration struct {
	Concurrency    int32 `yaml:"concurrency"`
	BatchFrequency int   `yaml:"batch-frequency"`
	BatchLimit     int   `yaml:"batch-limit"`

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

// Override
// 覆盖配置.
func (o *Configuration) Override(c *Configuration) *Configuration {
	if c.Concurrency > 0 {
		o.Concurrency = c.Concurrency
	}
	if c.BatchFrequency > 0 {
		o.BatchFrequency = c.BatchFrequency
	}
	if c.BatchLimit > 0 {
		o.BatchLimit = c.BatchLimit
	}

	if c.KeyLifetime > 0 {
		o.KeyLifetime = c.KeyLifetime
	}
	if c.KeyList != "" {
		o.KeyList = c.KeyList
	}
	if c.KeyPrefix != "" {
		o.KeyPrefix = c.KeyPrefix
	}

	if c.Network != "" {
		o.Network = c.Network
	}
	if c.Address != "" {
		o.Address = c.Address
	}
	if c.Password != "" {
		o.Password = c.Password
	}
	if c.Database > 0 {
		o.Database = c.Database
	}

	if c.MaxActive > 0 {
		o.MaxActive = c.MaxActive
	}
	if c.MaxIdle > 0 {
		o.MaxIdle = c.MaxIdle
	}
	if c.Timeout > 0 {
		o.Timeout = c.Timeout
	}
	if c.ReadTimeout > 0 {
		o.ReadTimeout = c.ReadTimeout
	}
	if c.WriteTimeout > 0 {
		o.WriteTimeout = c.WriteTimeout
	}
	if c.Wait != nil {
		o.Wait = c.Wait
	}

	return o
}

// 构造实例.
func (o *Configuration) init() *Configuration {
	o.Concurrency = defaultConcurrency
	o.BatchFrequency = defaultBatchFrequency
	o.BatchLimit = defaultBatchLimit

	o.KeyLifetime = defaultKeyLifetime
	o.KeyList = defaultKeyList
	o.KeyPrefix = defaultKeyPrefix

	o.Network = defaultNetwork
	o.Address = defaultAddress
	o.MaxActive = defaultMaxActive
	o.MaxIdle = defaultMaxIdle
	o.Timeout = defaultTimeout
	o.ReadTimeout = defaultReadTimeout
	o.WriteTimeout = defaultWriteTimeout
	o.Wait = &defaultPoolWait
	return o
}

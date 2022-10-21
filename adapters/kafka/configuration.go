// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package kafka

var Config *Configuration

const (
	defaultBatch              = 100
	defaultRetrySeconds       = 1
	defaultRetryTimes   int32 = 3

	defaultKafkaConnectRequests  = 5
	defaultKafkaConnectKeepAlive = 25
	defaultKafkaConnectTimeout   = 3
	defaultKafkaReadTimeout      = 2
	defaultKafkaWriteTime        = 10
	defaultKafkaFlushMessages    = 100
	defaultKafkaFlushSeconds     = 1
	defaultKafkaBroker           = "127.0.0.1:9092"
	defaultKafkaTopic            = "flog"
	defaultKafkaMaxMessageSize   = 2048
)

// Configuration
// 基础配置.
type Configuration struct {
	Debugger bool `yaml:"debugger"`

	KafkaConnectKeepAlive int      `yaml:"kafka-connect-keep-alive"`
	KafkaConnectRequests  int      `yaml:"kafka-connect-requests"`
	KafkaConnectTimeout   int      `yaml:"kafka-connect-timeout"`
	KafkaReadTimeout      int      `yaml:"kafka-read-timeout"`
	KafkaWriteTimeout     int      `yaml:"kafka-write-timeout"`
	KafkaFlushMessages    int      `yaml:"kafka-flush-messages"`
	KafkaFlushSeconds     int      `yaml:"kafka-flush-seconds"`
	KafkaMaxMessageSize   int      `yaml:"kafka-max-message-size"`
	KafkaBrokers          []string `yaml:"kafka-brokers"`
	KafkaTopic            string   `yaml:"kafka-topic"`

	RetrySeconds int   `yaml:"retry-seconds"`
	RetryTimes   int32 `yaml:"retry-times"`

	Batch       int   `yaml:"batch"`
	Concurrency int32 `yaml:"-"`
}

// Override
// 覆盖配置.
func (o *Configuration) Override(c *Configuration) *Configuration {
	o.Debugger = c.Debugger

	// Kafka 配置.
	if c.KafkaConnectKeepAlive > 0 {
		o.KafkaConnectKeepAlive = c.KafkaConnectKeepAlive
	}
	if c.KafkaConnectRequests > 0 {
		o.KafkaConnectRequests = c.KafkaConnectRequests
	}
	if c.KafkaConnectTimeout > 0 {
		o.KafkaConnectTimeout = c.KafkaConnectTimeout
	}
	if c.KafkaReadTimeout > 0 {
		o.KafkaReadTimeout = c.KafkaReadTimeout
	}
	if c.KafkaWriteTimeout > 0 {
		o.KafkaWriteTimeout = c.KafkaWriteTimeout
	}
	if c.KafkaFlushMessages > 0 {
		o.KafkaFlushMessages = c.KafkaFlushMessages
	}
	if c.KafkaFlushSeconds > 0 {
		o.KafkaFlushSeconds = c.KafkaFlushSeconds
	}
	if c.KafkaMaxMessageSize > 0 {
		o.KafkaMaxMessageSize = c.KafkaMaxMessageSize
	}
	if c.KafkaBrokers != nil && len(c.KafkaBrokers) > 0 {
		o.KafkaBrokers = c.KafkaBrokers
	}
	if c.KafkaTopic != "" {
		o.KafkaTopic = c.KafkaTopic
	}

	// Retry 配置.
	if c.RetrySeconds > 0 {
		o.RetrySeconds = c.RetrySeconds
	}
	if c.RetryTimes > 0 {
		o.RetryTimes = c.RetryTimes
	}

	// Basic 配置.
	if c.Batch > 0 {
		o.Batch = c.Batch
	}

	// Execution 属性.
	o.Concurrency = int32(o.KafkaConnectRequests)
	return o
}

// 构造实例.
func (o *Configuration) init() *Configuration {
	o.KafkaConnectKeepAlive = defaultKafkaConnectKeepAlive
	o.KafkaConnectRequests = defaultKafkaConnectRequests
	o.KafkaConnectTimeout = defaultKafkaConnectTimeout
	o.KafkaReadTimeout = defaultKafkaReadTimeout
	o.KafkaWriteTimeout = defaultKafkaWriteTime
	o.KafkaFlushMessages = defaultKafkaFlushMessages
	o.KafkaFlushSeconds = defaultKafkaFlushSeconds
	o.KafkaMaxMessageSize = defaultKafkaMaxMessageSize * 1024
	o.KafkaBrokers = []string{defaultKafkaBroker}
	o.KafkaTopic = defaultKafkaTopic

	o.RetrySeconds = defaultRetrySeconds
	o.RetryTimes = defaultRetryTimes

	o.Batch = defaultBatch
	o.Concurrency = int32(defaultKafkaConnectRequests)
	return o
}

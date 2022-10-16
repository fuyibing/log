// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package kafka

var Config *Configuration

const (
	defaultBatch   = 300
	defaultClean   = 1
	defaultBroker  = "127.0.0.1:9092"
	defaultTopic   = "flog"
	defaultTimeout = 3
)

// Configuration
// 基础配置.
type Configuration struct {
	Batch   int      `yaml:"batch"`
	Clean   int64    `yaml:"clean"`
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
	Timeout int      `yaml:"timeout"` // 连接超时时长
}

// Override
// 覆盖配置.
func (o *Configuration) Override(c *Configuration) *Configuration {
	// Basic 选项.

	if c.Batch > 0 {
		o.Batch = c.Batch
	}
	if c.Clean > 0 {
		o.Clean = c.Clean
	}

	// Kafka 选项.

	if c.Brokers != nil && len(c.Brokers) > 0 {
		o.Brokers = c.Brokers
	}
	if c.Topic != "" {
		o.Topic = c.Topic
	}
	if c.Timeout > 0 {
		o.Timeout = c.Timeout
	}
	return o
}

// 构造实例.
func (o *Configuration) init() *Configuration {
	o.Batch = defaultBatch
	o.Clean = defaultClean

	o.Brokers = []string{defaultBroker}
	o.Topic = defaultTopic
	o.Timeout = defaultTimeout

	return o
}

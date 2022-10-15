// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package kafka

var Config *Configuration

const (
	defaultConcurrency    int32 = 10
	defaultBatchFrequency       = 200
	defaultBatchLimit           = 100
	defaultBroker               = "127.0.0.1:9092"
	defaultTopic                = "flog"
)

// Configuration
// 基础配置.
type Configuration struct {
	Concurrency    int32 `yaml:"concurrency"`
	BatchFrequency int   `yaml:"batch-frequency"`
	BatchLimit     int   `yaml:"batch-limit"`

	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
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

	if c.Brokers != nil && len(c.Brokers) > 0 {
		o.Brokers = c.Brokers
	}

	if c.Topic != "" {
		o.Topic = c.Topic
	}

	return o
}

// 构造实例.
func (o *Configuration) init() *Configuration {
	o.Concurrency = defaultConcurrency
	o.BatchFrequency = defaultBatchFrequency
	o.BatchLimit = defaultBatchLimit
	o.Brokers = []string{defaultBroker}
	o.Topic = defaultTopic
	return o
}

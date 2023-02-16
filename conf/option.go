// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package conf

// Option
// interface for configuration fields configure.
type Option func(config *Configuration)

// WithAdapter bind adapter priority from left to right. If left
// executed error, select down.
//
// configured : Kafka, File.
// execution  : Return if logs send to kafka succeed, otherwise
//              write logs to disk file.
func WithAdapter(adapters ...Adapter) Option {
	return func(config *Configuration) {
		config.Adapters = adapters
	}
}

func WithLevel(level Level) Option {
	return func(config *Configuration) {
		config.Level = level
		config.updateStatus()
	}
}

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

// WithPrefix prepend specified string ahead of log text.
//
//   WithPrefix("MyLogger")
func WithPrefix(prefix string) Option {
	return func(config *Configuration) {
		config.Prefix = prefix
	}
}

// WithLevel specify log level.
//
//   WithLevel(conf.Debug)
//   WithLevel(conf.Info)
//   WithLevel(conf.Warn)
func WithLevel(level Level) Option {
	return func(config *Configuration) {
		config.Level = level
		config.updateStatus()
	}
}

// WithService set running service info.
//
//   WithService("[service=myapp]")
func WithService(service string) Option {
	return func(config *Configuration) {
		config.Service = service
	}
}

// WithTimeFormat specify format time string.
//
//   WithTimeFormat("2006-01-02 15:04:05.999")
func WithTimeFormat(s string) Option {
	return func(config *Configuration) {
		config.TimeFormat = s
	}
}

// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package conf

import (
	"path/filepath"
)

type (
	Option func(c *configuration)
)

// /////////////////////////////////////////////////////////////
// Basic: normal
// /////////////////////////////////////////////////////////////

func SetAdapter(s string) Option {
	return func(c *configuration) {
		c.Adapter = s
	}
}

func SetLevel(level Level) Option {
	return func(c *configuration) {
		c.Level = level

		// Compare level integer.
		i := level.Int()
		e := i > 0

		// Update switch status.
		c.fatalOn = e && i >= Fatal.Int()
		c.errorOn = e && i >= Error.Int()
		c.warnOn = e && i >= Warn.Int()
		c.infoOn = e && i >= Info.Int()
		c.debugOn = e && i >= Debug.Int()
	}
}

func SetPrefix(s string) Option {
	return func(c *configuration) {
		c.Prefix = s
	}
}

func SetServiceHost(s string) Option {
	return func(c *configuration) {
		c.ServiceHost = s
	}
}

func SetServiceName(s string) Option {
	return func(c *configuration) {
		c.ServiceName = s
	}
}

func SetServicePort(n int) Option {
	return func(c *configuration) {
		c.ServicePort = n
	}
}

func SetTimeFormat(s string) Option {
	return func(c *configuration) {
		c.TimeFormat = s
	}
}

// /////////////////////////////////////////////////////////////
// Basic: batch mode
// /////////////////////////////////////////////////////////////

func SetBatchConcurrency(n int32) Option {
	return func(c *configuration) {
		c.BatchConcurrency = n
	}
}

func SetBatchFrequency(n int) Option {
	return func(c *configuration) {
		c.BatchFrequency = n
	}
}

func SetBatchLimit(n int) Option {
	return func(c *configuration) {
		c.BatchLimit = n
	}
}

// /////////////////////////////////////////////////////////////
// Basic: open tracing
// /////////////////////////////////////////////////////////////

func SetSpanId(s string) Option {
	return func(c *configuration) {
		c.SpanId = s
	}
}

func SetParentSpanId(s string) Option {
	return func(c *configuration) {
		c.ParentSpanId = s
	}
}

func SetTraceId(s string) Option {
	return func(c *configuration) {
		c.TraceId = s
	}
}

func SetTraceVersion(s string) Option {
	return func(c *configuration) {
		c.TraceVersion = s
	}
}

// /////////////////////////////////////////////////////////////
// Adapter: File
// /////////////////////////////////////////////////////////////

func SetFileBasePath(s string) Option {
	return func(c *configuration) {
		c.File.BasePath, _ = filepath.Abs(s)
	}
}

func SetFileExtName(s string) Option {
	return func(c *configuration) {
		c.File.ExtName = s
	}
}

func SetFileFileName(s string) Option {
	return func(c *configuration) {
		c.File.FileName = s
	}
}

func SetFileSeparatorPath(s string) Option {
	return func(c *configuration) {
		c.File.SeparatorPath = s
	}
}

// /////////////////////////////////////////////////////////////
// Adapter: Term
// /////////////////////////////////////////////////////////////

func SetTermColor(b bool) Option {
	return func(c *configuration) {
		c.Term.Color = b
	}
}

// /////////////////////////////////////////////////////////////
// Adapter: Kafka
// /////////////////////////////////////////////////////////////

func SetKafkaAddress(addr ...string) Option {
	return func(c *configuration) {
		c.Kafka.Addresses = addr
	}
}

func SetKafkaTopic(topic string) Option {
	return func(c *configuration) {
		c.Kafka.Topic = topic
	}
}

// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package conf

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type (
	Configuration interface {
		GetAdapter() string
		GetAutoStart() bool
		GetBatchConcurrency() int32
		GetBatchLimit() int
		GetFile() FileConfiguration
		GetKafka() KafkaConfiguration
		GetLevel() Level
		GetParentSpanId() string
		GetPid() int
		GetPrefix() string
		GetServiceName() string
		GetSpanId() string
		GetTerm() TermConfiguration
		GetTimeFormat() string
		GetTraceId() string
		GetTraceVersion() string

		Set(options ...Option) Configuration

		DebugOn() bool
		ErrorOn() bool
		FatalOn() bool
		InfoOn() bool
		WarnOn() bool
	}

	configuration struct {
		Adapter     string `yaml:"adapter"`
		AutoStart   bool   `yaml:"auto-start"`
		Level       Level  `yaml:"level"`
		Prefix      string `yaml:"prefix"`
		ServiceName string `yaml:"service-name"`
		TimeFormat  string `yaml:"time-format"`

		// Basic: batch mode.

		BatchConcurrency int32 `yaml:"batch-concurrency"`
		BatchLimit       int   `yaml:"batch-limit"`

		// Base: open tracing.

		ParentSpanId string `yaml:"parent-span-id"`
		SpanId       string `yaml:"span-id"`
		TraceId      string `yaml:"trace-id"`
		TraceVersion string `yaml:"trace-version"`

		// Adapter: supports initialize.

		File  *fileConfiguration  `yaml:"file"`
		Kafka *kafkaConfiguration `yaml:"kafka"`
		Term  *termConfiguration  `yaml:"term"`

		// Dynamic fields.

		pid                                       int
		debugOn, infoOn, warnOn, errorOn, fatalOn bool
	}
)

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *configuration) Set(options ...Option) Configuration {
	for _, option := range options {
		option(o)
	}
	return o
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *configuration) init() *configuration {
	o.pid = os.Getpid()

	// File configurations.
	if o.File == nil {
		o.File = &fileConfiguration{}
	}

	o.File.parent = o

	// Kafka configurations..
	if o.Kafka == nil {
		o.Kafka = &kafkaConfiguration{}
	}

	o.File.parent = o

	// Term configurations.
	if o.Term == nil {
		o.Term = &termConfiguration{}
	}

	o.File.parent = o

	return o
}

func (o *configuration) initYaml() {
	for _, path := range []string{
		"config/log.yaml", "../config/log.yaml",
		"tmp/log.yaml", "../tmp/log.yaml",
	} {
		if buf, err := os.ReadFile(path); err == nil {
			if yaml.Unmarshal(buf, o) == nil {
				break
			}
		}
	}
}

func (o *configuration) initDefaults() {
	// Basic: normal.

	if o.Adapter == "" {
		o.Set(SetAdapter(DefaultAdapter))
	}
	if s := o.Level.String(); s != "" {
		o.Set(SetLevel(Level(strings.ToUpper(s))))
	}
	if o.Level.Int() == 0 {
		o.Set(SetLevel(DefaultLevel))
	}
	if o.TimeFormat == "" {
		o.Set(SetTimeFormat(DefaultTimeFormat))
	}

	// Basic: batch mode.
	if o.BatchConcurrency == 0 {
		o.Set(SetBatchConcurrency(DefaultBatchConcurrency))
	}
	if o.BatchLimit == 0 {
		o.Set(SetBatchLimit(DefaultBatchLimit))
	}

	// Basic: open tracing.

	if o.ParentSpanId == "" {
		o.Set(SetParentSpanId(DefaultParentSpanId))
	}
	if o.SpanId == "" {
		o.Set(SetSpanId(DefaultSpanId))
	}
	if o.TraceId == "" {
		o.Set(SetTraceId(DefaultTraceId))
	}
	if o.TraceVersion == "" {
		o.Set(SetTraceVersion(DefaultTraceVersion))
	}

	// Adapters default fields.

	o.File.initDefaults()
	o.Kafka.initDefaults()
	o.Term.initDefaults()
}

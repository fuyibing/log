// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

var (
	defaultAdapters = []Adapter{Term}
)

const (
	defaultLevel            = Debug
	defaultTimeFormat       = "2006-01-02 15:04:05.999"
	defaultBatchConcurrency = 1
	defaultBatchLimit       = 100
	defaultBatchDuration    = 300
)

type Configuration struct {
	AutoStart  bool      `yaml:"auto-start"`
	Adapters   []Adapter `yaml:"adapters"`
	Level      Level     `yaml:"level"`
	Prefix     string    `yaml:"prefix"`
	Service    string    `yaml:"service"`
	TimeFormat string    `yaml:"time-format"`

	BatchConcurrency int32 `yaml:"batch-concurrency"`
	BatchLimit       int   `yaml:"batch-limit"`
	BatchDuration    int   `yaml:"batch-duration"`

	File  *FileConfig  `yaml:"file"`
	Term  *TermConfig  `yaml:"term"`
	Kafka *KafkaConfig `yaml:"kafka"`

	debugOn, infoOn, warnOn, errorOn, fatalOn bool
}

// /////////////////////////////////////////////////////////////
// Exported methods
// /////////////////////////////////////////////////////////////

func (o *Configuration) Set(options ...Option) {
	for _, option := range options {
		option(o)
	}
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *Configuration) defaults() {
	// Batch mode.
	if o.BatchConcurrency == 0 {
		o.BatchConcurrency = defaultBatchConcurrency
	}
	if o.BatchDuration == 0 {
		o.BatchDuration = defaultBatchDuration
	}
	if o.BatchLimit == 0 {
		o.BatchLimit = defaultBatchLimit
	}

	opts := make([]Option, 0)

	// Config defaults.
	if len(o.Adapters) == 0 {
		opts = append(opts, WithAdapter(defaultAdapters...))
	}
	if o.TimeFormat == "" {
		opts = append(opts, WithTimeFormat(defaultTimeFormat))
	}
	if o.Level.Int() == 0 {
		opts = append(opts, WithLevel(defaultLevel))
	} else {
		o.updateStatus()
	}

	// Config: override.
	if len(opts) > 0 {
		o.Set(opts...)
	}

	// Adapter: file
	if o.File == nil {
		o.File = &FileConfig{}
	}
	o.File.defaults()

	// Adapter: Kafka
	if o.Kafka == nil {
		o.Kafka = &KafkaConfig{}
	}
	o.Kafka.defaults()

	// Adapter: terminal
	if o.Term == nil {
		o.Term = &TermConfig{}
	}
	o.Term.defaults()
}

func (o *Configuration) init() *Configuration {
	o.scan()
	o.defaults()
	return o
}

func (o *Configuration) scan() {
	for _, f := range []string{"config/log.yaml", "../config/log.yaml", "tmp/log.yaml", "../tmp/log.yaml"} {
		if buf, err := os.ReadFile(f); err == nil {
			if yaml.Unmarshal(buf, o) == nil {
				break
			}
		}
	}
}

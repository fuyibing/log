// author: wsfuyibing <websearch@163.com>
// date: 2021-02-10

// Package log manager.
package log

import (
	"sync"
)

// Adapter constants.
const (
	AdapterTerm Adapter = iota
	AdapterFile
	AdapterRedis
)

// Level constants.
const (
	LevelOff Level = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

// Default for package.
const (
	DefaultAdapter     = AdapterTerm
	DefaultLevel       = LevelDebug
	DefaultTimeFormat  = "2006-01-02 15:04:05.999999"
	DefaultSpanId      = "X-B3-Spanid"
	DefaultSpanVersion = "X-B3-Version"
	DefaultTraceId     = "X-B3-Traceid"
	OpenTracingContext = "OpenTracingContext"
)

// Global.
var (
	Config *configuration
	Logger LogInterface
	// Adapter names define.
	AdapterText = map[Adapter]string{
		AdapterTerm:  "TERM",
		AdapterFile:  "FILE",
		AdapterRedis: "REDIS",
	}
	// Adapter levels define.
	LevelText = map[Level]string{
		LevelOff:   "OFF",
		LevelError: "ERROR",
		LevelWarn:  "WARN",
		LevelInfo:  "INFO",
		LevelDebug: "DEBUG",
	}
)

// Adapter type.
type Adapter int

// Level type.
type Level int

// Handled when package initialized.
func init() {
	new(sync.Once).Do(func() {
		Config = new(configuration)
		Config.initialize()
		Logger = New()
		Logger.SetAdapter(Config.Adapter)
		Logger.SetLevel(Config.Level)
	})
}

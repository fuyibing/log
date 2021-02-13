// author: wsfuyibing <websearch@163.com>
// date: 2021-02-10

package log

// Log struct.
type logger struct {
	adapter  Adapter
	callback Callback
	level    Level
	debugOn  bool
	infoOn   bool
	warnOn   bool
	errorOn  bool
}

// Log callback for adapter.
type Callback func(LineInterface)

// Log interface.
type LogInterface interface {
	Debug(string)
	Debugf(string, ...interface{})
	Debugfc(interface{}, string, ...interface{})
	DebugOn() bool
	Info(string)
	Infof(string, ...interface{})
	Infofc(interface{}, string, ...interface{})
	InfoOn() bool
	Warn(string)
	Warnf(string, ...interface{})
	Warnfc(interface{}, string, ...interface{})
	WarnOn() bool
	Error(string)
	Errorf(string, ...interface{})
	Errorfc(interface{}, string, ...interface{})
	ErrorOn() bool
	SetAdapter(Adapter) LogInterface
	SetCallback(Callback) LogInterface
	SetLevel(Level) LogInterface
	log(ctx interface{}, level Level, text string, args ...interface{})
}

// Create log default instance.
func New() LogInterface {
	return new(logger)
}

// Create log instance.
func Default() LogInterface {
	o := New()
	o.SetAdapter(DefaultAdapter)
	o.SetLevel(DefaultLevel)
	return o
}

// Send debug log.
func (o *logger) Debug(text string) {
	if o.debugOn {
		o.log(nil, LevelDebug, text)
	}
}

// Send debug log.
func (o *logger) Debugf(text string, args ...interface{}) {
	if o.debugOn {
		o.log(nil, LevelDebug, text, args...)
	}
}

// Send debug log.
func (o *logger) Debugfc(ctx interface{}, text string, args ...interface{}) {
	if o.debugOn {
		o.log(ctx, LevelDebug, text, args...)
	}
}

// Send debug log status.
func (o *logger) DebugOn() bool { return o.debugOn }

// Send info log.
func (o *logger) Info(text string) {
	if o.infoOn {
		o.log(nil, LevelInfo, text)
	}
}

// Send info log.
func (o *logger) Infof(text string, args ...interface{}) {
	if o.infoOn {
		o.log(nil, LevelInfo, text, args...)
	}
}

// Send info log.
func (o *logger) Infofc(ctx interface{}, text string, args ...interface{}) {
	if o.infoOn {
		o.log(ctx, LevelInfo, text, args...)
	}
}

// Send info log status.
func (o *logger) InfoOn() bool { return o.debugOn }

// Send warn log.
func (o *logger) Warn(text string) {
	if o.warnOn {
		o.log(nil, LevelWarn, text)
	}
}

// Send warn log.
func (o *logger) Warnf(text string, args ...interface{}) {
	if o.warnOn {
		o.log(nil, LevelWarn, text, args...)
	}
}

// Send warn log.
func (o *logger) Warnfc(ctx interface{}, text string, args ...interface{}) {
	if o.warnOn {
		o.log(ctx, LevelWarn, text, args...)
	}
}

// Send warn log status.
func (o *logger) WarnOn() bool { return o.debugOn }

// Send error log.
func (o *logger) Error(text string) {
	if o.errorOn {
		o.log(nil, LevelError, text)
	}
}

// Send error log.
func (o *logger) Errorf(text string, args ...interface{}) {
	if o.errorOn {
		o.log(nil, LevelError, text, args...)
	}
}

// Send error log.
func (o *logger) Errorfc(ctx interface{}, text string, args ...interface{}) {
	if o.errorOn {
		o.log(ctx, LevelError, text, args...)
	}
}

// Send error log status.
func (o *logger) ErrorOn() bool { return o.debugOn }

// Set log adapter.
func (o *logger) SetAdapter(adapter Adapter) LogInterface {
	o.adapter = adapter
	switch o.adapter {
	case AdapterTerm:
		o.callback = NewAdapterTerm().Callback
	case AdapterFile:
		o.callback = NewAdapterFile().Callback
	case AdapterRedis:
		o.callback = NewAdapterRedis().Callback
	}
	return o
}

// Set log callback.
func (o *logger) SetCallback(callback Callback) LogInterface {
	o.callback = callback
	return o
}

// Set log level.
func (o *logger) SetLevel(level Level) LogInterface {
	o.level = level
	if o.level <= LevelOff {
		o.debugOn = false
		o.infoOn = false
		o.warnOn = false
		o.errorOn = false
	} else {
		o.debugOn = o.level >= LevelDebug
		o.infoOn = o.level >= LevelInfo
		o.warnOn = o.level >= LevelWarn
		o.errorOn = o.level >= LevelError
	}
	return o
}

// Send log to adapter.
func (o *logger) log(ctx interface{}, level Level, text string, args ...interface{}) {
	if o.callback == nil {
		panic("unknown adapter or callback not specified")
	}
	o.callback(NewLine(ctx, level, text, args...))
}

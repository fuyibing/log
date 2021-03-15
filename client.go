// author: wsfuyibing <websearch@163.com>
// date: 2021-03-14

package log

type client struct {
	config *configuration
}

func (o *client) Debug(text string) {
	if o.config.debugOn {
		o.log(nil, LevelDebug, text)
	}
}

func (o *client) Debugf(text string, args ...interface{}) {
	if o.config.debugOn {
		o.log(nil, LevelDebug, text, args...)
	}
}

func (o *client) Debugfc(ctx interface{}, text string, args ...interface{}) {
	if o.config.debugOn {
		o.log(ctx, LevelDebug, text, args...)
	}
}

func (o *client) Info(text string) {
	if o.config.infoOn {
		o.log(nil, LevelInfo, text)
	}
}

func (o *client) Infof(text string, args ...interface{}) {
	if o.config.infoOn {
		o.log(nil, LevelInfo, text, args...)
	}
}

func (o *client) Infofc(ctx interface{}, text string, args ...interface{}) {
	if o.config.infoOn {
		o.log(ctx, LevelInfo, text, args...)
	}
}

func (o *client) Warn(text string) {
	if o.config.warnOn {
		o.log(nil, LevelWarn, text)
	}
}

func (o *client) Warnf(text string, args ...interface{}) {
	if o.config.warnOn {
		o.log(nil, LevelWarn, text, args...)
	}
}

func (o *client) Warnfc(ctx interface{}, text string, args ...interface{}) {
	if o.config.warnOn {
		o.log(ctx, LevelWarn, text, args...)
	}
}

func (o *client) Error(text string) {
	if o.config.errorOn {
		o.log(nil, LevelError, text)
	}
}

func (o *client) Errorf(text string, args ...interface{}) {
	if o.config.errorOn {
		o.log(nil, LevelError, text, args...)
	}
}

func (o *client) Errorfc(ctx interface{}, text string, args ...interface{}) {
	if o.config.errorOn {
		o.log(ctx, LevelError, text, args...)
	}
}

func (o *client) log(ctx interface{}, level Level, text string, args ...interface{}) {
	if o.config.handler != nil {
		o.config.handler(NewLine(ctx, level, text, args...))
	}
}

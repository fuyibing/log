// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package log

import "context"

func Debug(text string) {
	if Config.debugOn {
		Client.Debug(text)
	}
}

func Debugf(text string, args ...interface{}) {
	if Config.debugOn {
		Client.Debugf(text, args...)
	}
}

func Debugfc(ctx context.Context, text string, args ...interface{}) {
	if Config.debugOn {
		Client.Debugfc(ctx, text, args...)
	}
}

func Info(text string) {
	if Config.infoOn {
		Client.Info(text)
	}
}

func Infof(text string, args ...interface{}) {
	if Config.infoOn {
		Client.Infof(text, args...)
	}
}

func Infofc(ctx context.Context, text string, args ...interface{}) {
	if Config.infoOn {
		Client.Infofc(ctx, text, args...)
	}
}

func Warn(text string) {
	if Config.warnOn {
		Client.Warn(text)
	}
}

func Warnf(text string, args ...interface{}) {
	if Config.warnOn {
		Client.Warnf(text, args...)
	}
}

func Warnfc(ctx context.Context, text string, args ...interface{}) {
	if Config.warnOn {
		Client.Warnfc(ctx, text, args...)
	}
}

func Error(text string) {
	if Config.errorOn {
		Client.Error(text)
	}
}

func Errorf(text string, args ...interface{}) {
	if Config.errorOn {
		Client.Errorf(text, args...)
	}
}

func Errorfc(ctx context.Context, text string, args ...interface{}) {
	if Config.errorOn {
		Client.Errorfc(ctx, text, args...)
	}
}

func Panic(text string) {
	if Config.errorOn {
		Client.Panic(text)
	}
}

func Panicf(text string, args ...interface{}) {
	if Config.errorOn {
		Client.Panicf(text, args...)
	}
}

func Panicfc(ctx context.Context, text string, args ...interface{}) {
	if Config.errorOn {
		Client.Panicfc(ctx, text, args...)
	}
}

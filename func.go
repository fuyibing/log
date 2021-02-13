// author: wsfuyibing <websearch@163.com>
// date: 2021-02-10

package log

// Send debug log.
func Debug(text string) {
	if Logger.DebugOn() {
		Logger.log(nil, LevelDebug, text)
	}
}

// Send debug log.
func Debugf(text string, args ...interface{}) {
	if Logger.DebugOn() {
		Logger.log(nil, LevelDebug, text, args...)
	}
}

// Send debug log.
func Debugfc(ctx interface{}, text string, args ...interface{}) {
	if Logger.DebugOn() {
		Logger.log(ctx, LevelDebug, text, args...)
	}
}

// Send info log.
func Info(text string) {
	if Logger.InfoOn() {
		Logger.log(nil, LevelInfo, text)
	}
}

// Send info log.
func Infof(text string, args ...interface{}) {
	if Logger.InfoOn() {
		Logger.log(nil, LevelInfo, text, args...)
	}
}

// Send info log.
func Infofc(ctx interface{}, text string, args ...interface{}) {
	if Logger.InfoOn() {
		Logger.log(ctx, LevelInfo, text, args...)
	}
}

// Send warn log.
func Warn(text string) {
	if Logger.WarnOn() {
		Logger.log(nil, LevelWarn, text)
	}
}

// Send warn log.
func Warnf(text string, args ...interface{}) {
	if Logger.WarnOn() {
		Logger.log(nil, LevelWarn, text, args...)
	}
}

// Send warn log.
func Warnfc(ctx interface{}, text string, args ...interface{}) {
	if Logger.WarnOn() {
		Logger.log(ctx, LevelWarn, text, args...)
	}
}

// Send error log.
func Error(text string) {
	if Logger.ErrorOn() {
		Logger.log(nil, LevelError, text)
	}
}

// Send error log.
func Errorf(text string, args ...interface{}) {
	if Logger.ErrorOn() {
		Logger.log(nil, LevelError, text, args...)
	}
}

// Send error log.
func Errorfc(ctx interface{}, text string, args ...interface{}) {
	if Logger.ErrorOn() {
		Logger.log(ctx, LevelError, text, args...)
	}
}

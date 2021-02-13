// author: wsfuyibing <websearch@163.com>
// date: 2021-02-13

package log

import (
	"fmt"

	"xorm.io/xorm/log"
)

// X ORM logger struct.
type XOrmLog struct{}

// Create X ORM logger instance.
func NewXOrmLog() *XOrmLog {
	return &XOrmLog{}
}

func (o *XOrmLog) Debugf(format string, v ...interface{}) {}
func (o *XOrmLog) Errorf(format string, v ...interface{}) {}
func (o *XOrmLog) Infof(format string, v ...interface{})  {}
func (o *XOrmLog) Warnf(format string, v ...interface{})  {}
func (o *XOrmLog) Level() log.LogLevel                    { return log.LOG_INFO }
func (o *XOrmLog) SetLevel(l log.LogLevel)                {}
func (o *XOrmLog) ShowSQL(show ...bool)                   {}
func (o *XOrmLog) IsShowSQL() bool                        { return true }
func (o *XOrmLog) BeforeSQL(c log.LogContext)             {}

// Send SQL to logger.
func (o *XOrmLog) AfterSQL(c log.LogContext) {
	// add INFO log.
	if Logger.InfoOn() {
		if c.Args != nil && len(c.Args) > 0 {
			Logger.log(c.Ctx, LevelInfo, fmt.Sprintf("[SQL][d=%f] %s - %v.", c.ExecuteTime.Seconds(), c.SQL, c.Args))
		} else {
			Logger.log(c.Ctx, LevelInfo, fmt.Sprintf("[SQL][d=%f] %s.", c.ExecuteTime.Seconds(), c.SQL))
		}
	}
	// add ERROR log.
	if c.Err != nil && Logger.ErrorOn() {
		Logger.log(c.Ctx, LevelError, fmt.Sprintf("[SQL] %s.", c.Err.Error()))
	}
}

// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package plugins

import (
	"fmt"

	xorm "xorm.io/xorm/log"

	"github.com/fuyibing/log"
)

type XOrm struct{}

func NewXOrm() *XOrm                                   { return &XOrm{} }
func (o *XOrm) Debugf(format string, v ...interface{}) {}
func (o *XOrm) Errorf(format string, v ...interface{}) {}
func (o *XOrm) Infof(format string, v ...interface{})  {}
func (o *XOrm) Warnf(format string, v ...interface{})  {}
func (o *XOrm) Level() xorm.LogLevel                   { return xorm.LOG_INFO }
func (o *XOrm) SetLevel(l xorm.LogLevel)               {}
func (o *XOrm) ShowSQL(show ...bool)                   {}
func (o *XOrm) IsShowSQL() bool                        { return true }
func (o *XOrm) BeforeSQL(c xorm.LogContext)            {}

// Send SQL to logger.
func (o *XOrm) AfterSQL(c xorm.LogContext) {
	// add INFO log.
	if log.Config.InfoOn() {
		if c.Args != nil && len(c.Args) > 0 {
			log.Client.Infofc(c.Ctx, fmt.Sprintf("[SQL][d=%f] %s - %v.", c.ExecuteTime.Seconds(), c.SQL, c.Args))
		} else {
			log.Client.Infofc(c.Ctx, fmt.Sprintf("[SQL][d=%f] %s.", c.ExecuteTime.Seconds(), c.SQL))
		}
	}
	// add ERROR log.
	if c.Err != nil && log.Config.ErrorOn() {
		log.Client.Errorfc(c.Ctx, fmt.Sprintf("[SQL] %s.", c.Err.Error()))
	}
}

// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package plugins

import (
	"fmt"

	"xorm.io/builder"
	xormLog "xorm.io/xorm/log"

	"github.com/fuyibing/log/v2"
)

type XOrm struct{}

func NewXOrm() *XOrm                                   { return &XOrm{} }
func (o *XOrm) Debugf(format string, v ...interface{}) {}
func (o *XOrm) Errorf(format string, v ...interface{}) {}
func (o *XOrm) Infof(format string, v ...interface{})  {}
func (o *XOrm) Warnf(format string, v ...interface{})  {}
func (o *XOrm) Level() xormLog.LogLevel                { return xormLog.LOG_INFO }
func (o *XOrm) SetLevel(l xormLog.LogLevel)            {}
func (o *XOrm) ShowSQL(show ...bool)                   {}
func (o *XOrm) IsShowSQL() bool                        { return true }
func (o *XOrm) BeforeSQL(c xormLog.LogContext)         {}

// Send SQL to logger.
func (o *XOrm) AfterSQL(c xormLog.LogContext) {
	// xorm session id
	var sId string
	v := c.Ctx.Value(xormLog.SessionIDKey)
	if key, ok := v.(string); ok {
		sId = key
	}
	// add INFO log.
	if log.Config.InfoOn() {
		if c.Args != nil && len(c.Args) > 0 {
			boundSql, _ := builder.ConvertToBoundSQL(c.SQL, c.Args)
			log.Client.Infofc(c.Ctx, fmt.Sprintf("[SQL=%s][d=%fs] %s - %v.", sId, c.ExecuteTime.Seconds(), boundSql, c.Args))
		} else {
			log.Client.Infofc(c.Ctx, fmt.Sprintf("[SQL=%s][d=%fs] %s.", sId, c.ExecuteTime.Seconds(), c.SQL))
		}
	}
	// add ERROR log.
	if c.Err != nil && log.Config.ErrorOn() {
		log.Client.Errorfc(c.Ctx, fmt.Sprintf("[SQL][sId=%s][dur=%fs] %s.", sId, c.ExecuteTime.Seconds(), c.Err.Error()))
	}
}

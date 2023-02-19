// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package main

import (
	"github.com/fuyibing/log/v8"
	"github.com/fuyibing/log/v8/adapters"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
	"regexp"
)

func init() {
	log.Config.Set(
		conf.SetAdapter(adapters.AdapterTerm),
		conf.SetLevel(conf.Debug),
		conf.SetTermColor(true),
	)

	regexDemo := regexp.MustCompile(`\s+demo\s+`)

	base.Parser.Set("demo", func(line *base.Line) {
		line.Text = regexDemo.ReplaceAllString(line.Text, " {MyAPP} ")
	})
}

func main() {
	// Close log client. Ensure that all data in the
	// memory queue are processed.
	defer log.Client.Close()

	log.Debug("example 1 debug")
	log.Info("[d=1.23] demo 2 info")
	log.Warn("[D=2.345] example 3 warn")
	log.Error("[Duration=3.456] example 4 error")
	log.Panic("example 5 panic")
}

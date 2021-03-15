// author: wsfuyibing <websearch@163.com>
// date: 2021-03-14

package log

import (
	"fmt"
)

type Adapter int

// func (c Adapter) Text() (text string) {
// 	switch c {
// 	case AdapterTerm:
// 		text = "TERM"
// 	case AdapterFile:
// 		text = "FILE"
// 	case AdapterRedis:
// 		text = "REDIS"
// 	}
// 	return
// }

const (
	AdapterOff Adapter = iota
	AdapterTerm
	AdapterFile
	AdapterRedis
)

var Adapters = map[Adapter]string{
	AdapterOff:   "OFF",
	AdapterTerm:  "TERM",
	AdapterFile:  "FILE",
	AdapterRedis: "REDIS",
}

type Handler func(*Line)

type adapterTerm struct{}

func newAdapterTerm() *adapterTerm { return &adapterTerm{} }

func (o *adapterTerm) handler(line *Line) {
	println(o.format(line))
}

// Format terminal content.
func (o *adapterTerm) format(line *Line) string {
	s := fmt.Sprintf("[%s][%5s]", line.GetTimeline(), line.GetLevel().Text())
	// Tracing.
	if line.HasTracing() {
		if method, route := line.GetTracing().GetRequestInfo(); method != "" {
			s += fmt.Sprintf("[v=%s][m=%s][u=%s]", line.GetTracing().Version(line.GetOffset()), method, route)
		} else {
			s += fmt.Sprintf("[v=%s]", line.GetTracing().Version(line.GetOffset()))
		}
	}
	// Content.
	return s + " " + line.String()
}

type adapterFile struct{}

func newAdapterFile() *adapterFile        { return &adapterFile{} }
func (o *adapterFile) handler(line *Line) {}

type adapterRedis struct{}

func newAdapterRedis() *adapterRedis       { return &adapterRedis{} }
func (o *adapterRedis) handler(line *Line) {}

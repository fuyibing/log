// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package base

import (
	"regexp"
	"strconv"
	"sync"
)

type (
	ParserHandler func(line *Line)

	ParserManager interface {
		Parse(line *Line)
		Set(key string, handler ParserHandler) ParserManager
		Unset(key string) ParserManager
	}

	parser struct {
		handlers map[string]ParserHandler
		mu       *sync.RWMutex
	}
)

func (o *parser) Parse(line *Line) {
	for _, handler := range func() map[string]ParserHandler {
		o.mu.RLock()
		defer o.mu.RUnlock()
		return o.handlers
	}() {
		if handler != nil {
			handler(line)
		}
	}
}

func (o *parser) Set(key string, handler ParserHandler) ParserManager {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.handlers[key] = handler
	return o
}

func (o *parser) Unset(key string) ParserManager {
	o.mu.Lock()
	defer o.mu.Unlock()
	delete(o.handlers, key)
	return o
}

func (o *parser) init() *parser {
	o.handlers = make(map[string]ParserHandler)
	o.mu = &sync.RWMutex{}
	o.initDefaults()
	return o
}

func (o *parser) initDefaults() {
	var (
		regexDuration = regexp.MustCompile(`(?i)\[(d|dur|duration)=(\d+\.?\d*)\]\s*`)
	)

	o.Set("duration", func(line *Line) {
		// Parse duration keyword from text.
		//
		//   - [d=1.23]
		//   - [dur=1.23]
		//   - [duration=1.23]
		if m := regexDuration.FindStringSubmatch(line.Text); len(m) == 3 {
			if f, fe := strconv.ParseFloat(m[2], 64); fe == nil {
				line.Duration = f
				line.Text = regexDuration.ReplaceAllString(line.Text, "")
			}
		}
	})
}

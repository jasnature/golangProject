package appenders

import (
	"GoBLog/base"
)

//can support mutiple type appender.
type multipleOutputAppender struct {
	Appender

	appenderList []Appender

	syncChan chan chanMsg
}

type chanMsg struct {
	level   base.LogLevel
	message string
	args    []interface{}
}

func NewMultipleAppender(maxQueue int, appenders ...Appender) Appender {
	if maxQueue <= 0 || maxQueue >= 1000 {
		maxQueue = 100
	}
	obj := &multipleOutputAppender{
		appenderList: appenders,
		syncChan:     make(chan chanMsg, maxQueue),
	}
	go obj.processWriteString()
	return obj
}

func (this *multipleOutputAppender) processWriteString() {
	for data := range this.syncChan {
		for _, appender := range this.appenderList {
			appender.WriteString(data.level, data.message, data.args...)
		}
	}
}

func (this *multipleOutputAppender) WriteString(level base.LogLevel, message string, args ...interface{}) {
	this.syncChan <- chanMsg{
		level,
		message,
		args,
	}
}
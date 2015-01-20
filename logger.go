package gol

import (
	"fmt"
	"time"

	"github.com/kavkaz/gore"
)

func (logger *Logger) Error(args ...interface{}) {
	if logger.Level >= ERROR {
		logger.process(&Record{Time: time.Now(), Level: ERROR, Body: fmt.Sprint(args...)})
	}
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	if logger.Level >= ERROR {
		logger.process(&Record{Time: time.Now(), Level: ERROR, Body: fmt.Sprintf(format, args...)})
	}
}

func (logger *Logger) ErrorE(err error) {
	if logger.Level >= ERROR {
		msg, context := errorInfo(err)
		logger.process(&Record{Time: time.Now(), Level: ERROR, Body: msg, Context: context})
	}
}

func errorInfo(err error) (string, []string) {
	var msg string
	context := []string{}

	gerr, ok := err.(*gore.Err)
	if ok {
		msg = fmt.Sprintf("[%s %s] %s", gerr.Caller.ShortFileName(), gerr.Caller.ShortFuncName(), gerr.Error())
		if gerr.Context != nil {
			for i, c := range gerr.Context {
				context = append(context, fmt.Sprintf("#%d [%s %s] %s", i, c.Caller.ShortFileName(), c.Caller.ShortFuncName(), c.String()))
			}
		}
	} else {
		msg = err.Error()
	}

	return msg, context
}

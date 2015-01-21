package gol

import (
	"fmt"
	"os"

	"github.com/kavkaz/gore"
)

// Panic
func (logger *Logger) Panic(args ...interface{}) {
	logger.Log(3, PANIC, fmt.Sprint(args...), nil)
	panic(fmt.Sprint(args...))
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.Log(3, PANIC, fmt.Sprintf(format, args...), nil)
	panic(fmt.Sprintf(format, args...))
}

func (logger *Logger) PanicE(err error) {
	msg, context := errorInfo(err)
	logger.Log(3, PANIC, msg, context)
	panic(fmt.Sprint(err))
}

// Fatal
func (logger *Logger) Fatal(args ...interface{}) {
	logger.Log(3, FATAL, fmt.Sprint(args...), nil)
	os.Exit(1)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.Log(3, FATAL, fmt.Sprintf(format, args...), nil)
	os.Exit(1)
}

func (logger *Logger) FatalE(err error) {
	msg, context := errorInfo(err)
	logger.Log(3, FATAL, msg, context)
	os.Exit(1)
}

// Error
func (logger *Logger) Error(args ...interface{}) {
	logger.Log(3, ERROR, fmt.Sprint(args...), nil)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.Log(3, ERROR, fmt.Sprintf(format, args...), nil)
}

func (logger *Logger) ErrorE(err error) {
	msg, context := errorInfo(err)
	logger.Log(3, ERROR, msg, context)
}

// Warning
func (logger *Logger) Warn(args ...interface{}) {
	logger.Log(3, WARN, fmt.Sprint(args...), nil)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.Log(3, WARN, fmt.Sprintf(format, args...), nil)
}

func (logger *Logger) WarnE(err error) {
	msg, context := errorInfo(err)
	logger.Log(3, WARN, msg, context)
}

// Info
func (logger *Logger) Info(args ...interface{}) {
	logger.Log(3, INFO, fmt.Sprint(args...), nil)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.Log(3, INFO, fmt.Sprintf(format, args...), nil)
}

func (logger *Logger) InfoE(err error) {
	msg, context := errorInfo(err)
	logger.Log(3, INFO, msg, context)
}

// Debug
func (logger *Logger) Debug(args ...interface{}) {
	logger.Log(3, DEBUG, fmt.Sprint(args...), nil)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.Log(3, DEBUG, fmt.Sprintf(format, args...), nil)
}

func (logger *Logger) DebugE(err error) {
	msg, context := errorInfo(err)
	logger.Log(3, DEBUG, msg, context)
}
func errorInfo(err error) (string, []string) {
	var msg string
	var context []string

	gerr, ok := err.(*gore.Err)
	if ok {
		msg = fmt.Sprintf("[%s %s] %s", gerr.Caller.ShortFileName(), gerr.Caller.ShortFuncName(), gerr.Error())
		if gerr.Context != nil {
			context = []string{}
			for i, c := range gerr.Context {
				context = append(context, fmt.Sprintf("#%d [%s %s] %s", i, c.Caller.ShortFileName(), c.Caller.ShortFuncName(), c.String()))
			}
		}
	} else {
		msg = err.Error()
	}

	return msg, context
}

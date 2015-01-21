package gol

import (
	"fmt"
	"os"
	"time"

	"github.com/kavkaz/gore"
)

// Panic
func (logger *Logger) Panic(args ...interface{}) {
	if logger.level >= PANIC {
		logger.process(&Record{Time: time.Now(), Level: PANIC, Body: fmt.Sprint(args...)})
	}
	panic(fmt.Sprint(args...))
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	if logger.level >= PANIC {
		logger.process(&Record{Time: time.Now(), Level: PANIC, Body: fmt.Sprintf(format, args...)})
	}
	panic(fmt.Sprintf(format, args...))
}

func (logger *Logger) PanicE(err error) {
	if logger.level >= PANIC {
		msg, context := errorInfo(err)
		logger.process(&Record{Time: time.Now(), Level: PANIC, Body: msg, Context: context})
	}
	panic(fmt.Sprint(err))
}

// Fatal
func (logger *Logger) Fatal(args ...interface{}) {
	if logger.level >= FATAL {
		logger.process(&Record{Time: time.Now(), Level: FATAL, Body: fmt.Sprint(args...)})
	}
	os.Exit(1)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	if logger.level >= FATAL {
		logger.process(&Record{Time: time.Now(), Level: FATAL, Body: fmt.Sprintf(format, args...)})
	}
	os.Exit(1)
}

func (logger *Logger) FatalE(err error) {
	if logger.level >= FATAL {
		msg, context := errorInfo(err)
		logger.process(&Record{Time: time.Now(), Level: FATAL, Body: msg, Context: context})
	}
	os.Exit(1)
}

// Error
func (logger *Logger) Error(args ...interface{}) {
	if logger.level >= ERROR {
		logger.process(&Record{Time: time.Now(), Level: ERROR, Body: fmt.Sprint(args...)})
	}
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	if logger.level >= ERROR {
		logger.process(&Record{Time: time.Now(), Level: ERROR, Body: fmt.Sprintf(format, args...)})
	}
}

func (logger *Logger) ErrorE(err error) {
	if logger.level >= ERROR {
		msg, context := errorInfo(err)
		logger.process(&Record{Time: time.Now(), Level: ERROR, Body: msg, Context: context})
	}
}

// Warning
func (logger *Logger) Warn(args ...interface{}) {
	if logger.level >= WARN {
		logger.process(&Record{Time: time.Now(), Level: WARN, Body: fmt.Sprint(args...)})
	}
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	if logger.level >= WARN {
		logger.process(&Record{Time: time.Now(), Level: WARN, Body: fmt.Sprintf(format, args...)})
	}
}

func (logger *Logger) WarnE(err error) {
	if logger.level >= WARN {
		msg, context := errorInfo(err)
		logger.process(&Record{Time: time.Now(), Level: WARN, Body: msg, Context: context})
	}
}

// Info
func (logger *Logger) Info(args ...interface{}) {
	if logger.level >= INFO {
		logger.process(&Record{Time: time.Now(), Level: INFO, Body: fmt.Sprint(args...)})
	}
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	if logger.level >= INFO {
		logger.process(&Record{Time: time.Now(), Level: INFO, Body: fmt.Sprintf(format, args...)})
	}
}

func (logger *Logger) InfoE(err error) {
	if logger.level >= INFO {
		msg, context := errorInfo(err)
		logger.process(&Record{Time: time.Now(), Level: INFO, Body: msg, Context: context})
	}
}

// Debug
func (logger *Logger) Debug(args ...interface{}) {
	if logger.level >= DEBUG {
		logger.process(&Record{Time: time.Now(), Level: DEBUG, Body: fmt.Sprint(args...)})
	}
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	if logger.level >= DEBUG {
		logger.process(&Record{Time: time.Now(), Level: DEBUG, Body: fmt.Sprintf(format, args...)})
	}
}

func (logger *Logger) DebugE(err error) {
	if logger.level >= DEBUG {
		msg, context := errorInfo(err)
		logger.process(&Record{Time: time.Now(), Level: DEBUG, Body: msg, Context: context})
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

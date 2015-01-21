package gol

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var std = New()

func Output(buf *bytes.Buffer) {
	std.Output(buf)
}

func SetLevel(lvl Level) {
	std.SetLevel(lvl)
}

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func SetView(v Viewer) {
	std.SetView(v)
}

// Panic
func Panic(args ...interface{}) {
	std.Log(3, PANIC, fmt.Sprint(args...), nil)
	panic(fmt.Sprint(args...))
}

func Panicf(format string, args ...interface{}) {
	std.Log(3, PANIC, fmt.Sprintf(format, args...), nil)
	panic(fmt.Sprintf(format, args...))
}

func PanicE(err error) {
	msg, context := errorInfo(err)
	std.Log(3, PANIC, msg, context)
	panic(fmt.Sprint(err))
}

// Fatal
func Fatal(args ...interface{}) {
	std.Log(3, FATAL, fmt.Sprint(args...), nil)
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	std.Log(3, FATAL, fmt.Sprintf(format, args...), nil)
	os.Exit(1)
}

func FatalE(err error) {
	msg, context := errorInfo(err)
	std.Log(3, FATAL, msg, context)
	os.Exit(1)
}

// Error
func Error(args ...interface{}) {
	std.Log(3, ERROR, fmt.Sprint(args...), nil)
}

func Errorf(format string, args ...interface{}) {
	std.Log(3, ERROR, fmt.Sprintf(format, args...), nil)
}

func ErrorE(err error) {
	msg, context := errorInfo(err)
	std.Log(3, ERROR, msg, context)
}

// Warning
func Warn(args ...interface{}) {
	std.Log(3, WARN, fmt.Sprint(args...), nil)
}

func Warnf(format string, args ...interface{}) {
	std.Log(3, WARN, fmt.Sprintf(format, args...), nil)
}

func WarnE(err error) {
	msg, context := errorInfo(err)
	std.Log(3, WARN, msg, context)
}

// Info
func Info(args ...interface{}) {
	std.Log(3, INFO, fmt.Sprint(args...), nil)
}

func Infof(format string, args ...interface{}) {
	std.Log(3, INFO, fmt.Sprintf(format, args...), nil)
}

func InfoE(err error) {
	msg, context := errorInfo(err)
	std.Log(3, INFO, msg, context)
}

// Debug
func Debug(args ...interface{}) {
	std.Log(3, DEBUG, fmt.Sprint(args...), nil)
}

func Debugf(format string, args ...interface{}) {
	std.Log(3, DEBUG, fmt.Sprintf(format, args...), nil)
}

func DebugE(err error) {
	msg, context := errorInfo(err)
	std.Log(3, DEBUG, msg, context)
}

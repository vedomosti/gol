package gol

import (
	"io"
	"os"
	"sync"
)

type Level uint8

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

var LevelsString = map[Level]string{
	PanicLevel: "panic",
	FatalLevel: "fatal",
	ErrorLevel: "error",
	WarnLevel:  "warning",
	InfoLevel:  "info",
	DebugLevel: "debug",
}

func (level Level) String() string {
	return LevelsString[level]
}

type EncodeFormat uint8

const (
	TextEncodeFormat EncodeFormat = iota
	JsonEncodeFormat
)

type M map[string]string

type Logger struct {
	mu           sync.Mutex
	Level        Level
	EncodeFormat EncodeFormat
	Out          io.Writer
}

func New() *Logger {
	return &Logger{
		Level:        InfoLevel,
		EncodeFormat: TextEncodeFormat,
		Out:          os.Stdout,
	}
}

func (logger *Logger) Error() {

}

func (logger *Logger) Output() {

}

type Record struct {
	Logger *Logger
}

func NewRecord(logger *Logger) *Record {
	return &Record{
		Logger: logger,
	}
}

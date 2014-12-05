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
type F []string

const (
	TextEncodeFormat EncodeFormat = iota
	JsonEncodeFormat
)

type Formatter struct {
	Type   EncodeFormat
	Fields F
}

type M map[string]string

type Logger struct {
	mu        sync.Mutex
	Level     Level
	Formatter Formatter
	Out       io.Writer
}

func New() *Logger {
	return &Logger{
		Level:     InfoLevel,
		Formatter: Formatter{Type: TextEncodeFormat},
		Out:       os.Stdout,
	}
}

// todo: simple way
func (logger *Logger) SetFormatter(etype EncodeFormat, fields F) {
	logger.Formatter.Type = etype
	logger.Formatter.Fields = fields
}

func (logger *Logger) Error() {}

func (logger *Logger) Output() {}

type Record struct {
	Logger *Logger
}

func NewRecord(logger *Logger) *Record {
	return &Record{
		Logger: logger,
	}
}

func (rec *Record) String() string {
	return "fire"
}

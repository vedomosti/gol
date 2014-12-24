package gol

import (
	"io"
	"os"
	"sync"
)

type Level uint8

const (
	PANIC Level = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
)

var levelsString = []string{
	"PANIC",
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
}

func (level Level) String() string {
	return levelsString[level]
}

type EncodeFormat uint8

const (
	TEXT EncodeFormat = iota
	JSON
	PRETTY
)

type Logger struct {
	mu     sync.Mutex
	Level  Level
	Format EncodeFormat
	Out    io.Writer
}

func New() *Logger {
	return &Logger{
		Level:  INFO,
		Format: PRETTY,
		Out:    os.Stderr,
	}
}

func (logger *Logger) Error(args ...interface{}) {
	logger.process(ERROR, &Record{Body: args})
}

func (logger *Logger) process(level Level, record IRecord) {
	if logger.Level >= level {
		logger.Out.Write(record.Bytes())
	}
}

type IRecord interface {
	Bytes() []byte
}

package gol

import (
	"io"
	"os"
	"sync"
	"syscall"
	"time"
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
	level  Level
	format EncodeFormat
	out    io.Writer
}

func New() *Logger {
	return &Logger{
		level:  INFO,
		format: PRETTY,
		out:    os.Stderr,
	}
}

func (logger *Logger) Level() Level {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.level
}

func (logger *Logger) Format() EncodeFormat {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.format
}

func (logger *Logger) SetOutput(out io.Writer) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.out = out
}

func (logger *Logger) ignore(level Level) bool {
	return level > logger.Level()
}

func (logger *Logger) receive(level Level, args ...interface{}) {
	now := time.Now()
	if logger.ignore(level) {
		return
	}

	rbytes := &Record{
		Format: logger.Format(),
		Pid:    syscall.Getpid(),
		Time:   now,
		Level:  level,
		Caller: newCaller(2),
		Body:   args,
	}.Bytes()

	logger.mu.Lock()
	defer logger.mu.Unlock()

	logger.out.Write(rbytes)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.receive(ERROR, args...)
}

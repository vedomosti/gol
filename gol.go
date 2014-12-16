package gol

import (
	"io"
	"os"
	"sync"
	"time"
)

type LevelType uint8

const (
	PANIC LevelType = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
)

var LevelsString = map[LevelType]string{
	PANIC: "PANIC",
	FATAL: "FATAL",
	ERROR: "ERROR",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBUG",
}

func (level LevelType) String() string {
	return LevelsString[level]
}

type EncodeFormat uint8

const (
	TEXT EncodeFormat = iota
	JSON
)

type Logger struct {
	mu     sync.Mutex
	Level  Level
	Format EncodeFormat
	Out    io.Writer
}

func New() *Logger {
	return &Logger{
		Level:  InfoLevel,
		Format: EncodeFormat,
		Out:    os.Stdout,
	}
}

func (logger *Logger) Output() {}

type Record interface {
	ToString()
	ToJson()
}

type RecordHttpRequest struct {
	Ip     string
	Method string
	Url    string
}

type RecordHttpResponse struct {
	Status   int
	Url      string
	Duration time.Duration
}

type RecordSql struct {
	Query    string
	Params   interface{}
	Duration time.Duration
}

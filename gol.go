package gol

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"syscall"
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
	Level  LevelType
	Format EncodeFormat
	Out    io.Writer
}

func New() *Logger {
	return &Logger{
		Level:  INFO,
		Format: TEXT,
		Out:    os.Stdout,
	}
}

func (logger *Logger) Output(r *Record) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	r.Time = time.Now()
	r.Pid = syscall.Getpid()

	buf := r.Bytes()
	buf.WriteTo(logger.Out)
}

func (logger *Logger) Error(d interface{}) {
	logger.Output(&Record{Level: ERROR, Body: d})
}

type IRecord interface {
	String()
	Json()
}

type Record struct {
	Pid   int
	Time  time.Time
	Level LevelType
	Body  interface{}
}

func (rec *Record) Bytes() (buf bytes.Buffer) {
	buf.WriteString(fmt.Sprintf("%5d | ", syscall.Getpid()))
	buf.WriteString(rec.Time.Format("2006/01/02 15:04:05"))
	buf.WriteString("\n")

	return buf
}

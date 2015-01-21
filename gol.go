package gol

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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

type Viewer func(*Record) (*bytes.Buffer, error)

type Logger struct {
	mu    sync.Mutex
	level Level
	view  Viewer
	out   io.Writer
}

func New() *Logger {
	return &Logger{
		level: INFO,
		view:  ViewText,
		out:   os.Stderr,
	}
}

func (logger *Logger) Output(buf *bytes.Buffer) {
	logger.mu.Lock()
	logger.out.Write(buf.Bytes())
	logger.mu.Unlock()
}

func (logger *Logger) Level() Level {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.level
}

func (logger *Logger) SetLevel(lvl Level) {
	logger.mu.Lock()
	logger.level = lvl
	logger.mu.Unlock()
}

func (logger *Logger) SetOutput(w io.Writer) {
	logger.mu.Lock()
	logger.out = w
	logger.mu.Unlock()
}

func (logger *Logger) SetView(v Viewer) {
	logger.mu.Lock()
	logger.view = v
	logger.mu.Unlock()
}

func (logger *Logger) Log(calldepth int, level Level, msg string, context []string) {
	if logger.Level() >= level {
		logger.process(calldepth, &Record{Time: time.Now(), Level: level, Body: msg, Context: context})
	}
}

func (logger *Logger) process(calldepth int, record *Record) {
	record.Pid = syscall.Getpid()
	record.Caller = NewCaller(calldepth)

	buf, err := logger.view(record)
	if err != nil {
		logger.mu.Lock()
		fmt.Fprintf(logger.out, "Error generate view of record: %v\n", err)
		logger.mu.Unlock()
	}

	logger.Output(buf)
}

type Record struct {
	Time    time.Time
	Level   Level
	Pid     int
	Caller  *Caller
	Body    string
	Context []string
}

// Caller struct store info about caller
type Caller struct {
	FuncName string
	FileName string
	Line     int
}

func NewCaller(lvl int) *Caller {
	pc, fn, line, _ := runtime.Caller(lvl + 1)
	return &Caller{
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fn,
		Line:     line,
	}
}

func (caller *Caller) ShortFileName() string {
	return filepath.Base(caller.FileName) + ":" + strconv.Itoa(caller.Line)
}

func (caller *Caller) ShortFuncName() string {
	return filepath.Base(caller.FuncName)
}

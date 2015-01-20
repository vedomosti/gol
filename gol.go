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

type Logger struct {
	mu    sync.Mutex
	Level Level
	View  func(*Record) (*bytes.Buffer, error)
	out   io.Writer
}

func New() *Logger {
	return &Logger{
		Level: INFO,
		View:  ViewText,
		out:   os.Stderr,
	}
}

func (logger *Logger) process(record *Record) {
	record.Pid = syscall.Getpid()
	record.Caller = NewCaller(2)

	buf, err := logger.View(record)
	if err != nil {
		logger.mu.Lock()
		fmt.Fprintf(logger.out, "Error generate view of record: %v\n", err)
		logger.mu.Unlock()
	}

	logger.out.Write(buf.Bytes())
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

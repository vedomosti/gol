package gol

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
		out:   os.Stderr,
	}
}

func (logger *Logger) Error(args ...interface{}) {
	if logger.Level >= ERROR {
		logger.process(&Record{Time: time.Now(), Level: ERROR, Body: fmt.Sprint(args)})
	}
}

// todo ErrorE(err gore.Err)
// todo ErrorF

func (logger *Logger) process(record *Record) {
	record.Pid = syscall.Getpid()

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
	Body    string
	Context []string
}

func ViewText(record *Record) (*bytes.Buffer, error) {
	delimeter := []byte(" ")
	buf := bytes.NewBufferString(strconv.Itoa(record.Pid))
	buf.Write(delimeter)
	offset := buf.Len()
	buf.WriteString(record.Time.Format("2006/01/02 15:04:05"))
	buf.Write(delimeter)
	buf.WriteString(record.Level.String())
	buf.Write(delimeter)
	buf.WriteString(record.Body)
	if record.Context != nil {
		for _, msg := range record.Context {
			buf.WriteString("\n")
			buf.WriteString(strings.Repeat(" ", offset))
			buf.WriteString(msg)
		}
	}
	return buf, nil
}

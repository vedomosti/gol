package gol

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strconv"
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
	PRETTY
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
		Format: PRETTY,
		Out:    os.Stdout,
	}
}

// Record struct represent current record for logging
type Record struct {
	Level  LevelType
	Caller *Caller
	Body   interface{}
}

func (record *Record) PidString() string {
	return strconv.Itoa(syscall.Getpid())
}

func (record *Record) TimeString() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func (record *Record) Text() []byte {
	var buf bytes.Buffer
	delimeter := []byte{" "}

	buf.WriteString(record.PidString())
	buf.Write(delimeter)
	buf.WriteString(record.TimeString())
	buf.Write(delimeter)
	buf.WriteString(record.Level.String())
	buf.Write(delimeter)
	buf.WriteString(record.Caller.ShortFileName())
	buf.Write(delimeter)
	buf.WriteString(record.Caller.ShortFuncName())

	return buf.Bytes()
}

func (record *Record) Pretty() []byte {
	return []byte{}
}

func (record *Record) Json() []byte {
	return []byte{}
}

func (record *Record) Bytes(format EncodeFormat) []byte {
	switch format {
	default:
		return record.Text()
	case PRETTY:
		return record.Pretty()
	case JSON:
		return record.Json()
	}

	return nil
}

// Caller struct store info about caller
type Caller struct {
	FuncName string
	FileName string
	Line     int
}

func (caller *Caller) ShortFileName() string {
	return filepath.Base(caller.FileName) + ":" + strconv.Itoa(caller.Line)
}

func (caller *Caller) ShortFuncName() string {
	return filepath.Base(caller.FuncName)
}

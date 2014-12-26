package gol

import (
	"bytes"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

// Record struct represent current record for logging
type Record struct {
	Format EncodeFormat
	Pid    int
	Time   time.Time
	Level  Level
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
	delimeter := []byte(" ")

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

func (record *Record) Bytes() []byte {
	switch record.Format {
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

func newCaller(lvl int) *Caller {
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

package gol

import (
	"bytes"
	"strconv"
	"strings"
)

func ViewText(record *Record) (*bytes.Buffer, error) {
	delimeter := []byte(" ")
	buf := bytes.NewBufferString(strconv.Itoa(record.Pid))
	buf.Write(delimeter)
	offset := buf.Len()
	buf.WriteString(record.Time.Format("2006/01/02 15:04:05"))
	buf.Write(delimeter)
	buf.WriteString(record.Level.String())
	buf.Write(delimeter)

	if record.Caller != nil {
		buf.WriteString(record.Caller.ShortFileName())
		buf.Write(delimeter)
		buf.WriteString(record.Caller.ShortFuncName())
		buf.Write(delimeter)
	}

	buf.WriteString(record.Body)

	if record.Context != nil {
		for _, msg := range record.Context {
			buf.WriteString("\n")
			buf.WriteString(strings.Repeat(" ", offset))
			buf.WriteString(msg)
		}
	}
	buf.WriteString("\n")

	return buf, nil
}

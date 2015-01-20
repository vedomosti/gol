package gol

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLevelString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("FATAL", FATAL.String())
}

func TestViewText(t *testing.T) {
	assert := assert.New(t)
	r := &Record{Time: time.Now(), Pid: 123, Level: ERROR, Body: "hello"}
	r.Context = []string{"line1", "line2"}
	buf, _ := ViewText(r)
	assert.Contains(buf.String(), "hello")
	t.Log(buf.String())
}

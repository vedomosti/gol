package gol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("FATAL", FATAL.String())
}

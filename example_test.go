package gol

import (
	"bytes"
	"os"
)

func ExampleLogger_process() {
	logger := New()
	logger.Out = os.Stdout

	var buf bytes.Buffer
	buf.WriteString("hello")

	logger.process(ERROR, &buf)
	// Output:
	// hello
}

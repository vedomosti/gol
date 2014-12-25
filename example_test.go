package gol

import "os"

func ExampleLogger_process() {
	logger := New()
	logger.SetOutput(os.Stdout)

	logger.receive(ERROR, "hello")
	// Output:
	// hello
}

package main

import (
	"errors"

	"github.com/vedomosti/gol"
	"github.com/vedomosti/gore"
)

func main() {
	logger := gol.New()
	logger.Error("Hello worold")
	logger.Errorf("Some test: %v", map[string]int{"one": 1, "two": 2})

	err := gore.Newf("Hello %s", "world")
	gore.Append(err, "Foo bar")
	gore.Appendf(err, "Context %s", "info")
	logger.ErrorE(err)

	logger.ErrorE(errors.New("classic work!"))
}

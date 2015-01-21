package main

import "github.com/kavkaz/gol"

func main() {
	logger := gol.New()
	logger.Debug("Hidden Hello worold")
	logger.SetLevel(gol.DEBUG)
	logger.Debugf("Foo: %s", "Bar")
	logger.Info("notice message")
	logger.Warn("warning!!!")

}

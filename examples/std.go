package main

import "github.com/vedomosti/gol"

func main() {
	gol.Error("Hello world!")
	gol.SetLevel(gol.DEBUG)
	gol.Debugf("show: %v", map[string]int{"ont": 1})
	gol.Panic("opa opa")
}

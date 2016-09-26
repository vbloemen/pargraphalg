package main

import (
	"github.com/vbloemen/pargraphalg/graph"
)

func main() {

	g := graph.TestGraph2()

	g.PrintDOT()
}

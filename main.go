package main

import (
	"github.com/vbloemen/pargraphalg/alg"
	"github.com/vbloemen/pargraphalg/graph"
)

func main() {

	g := graph.TestGraph2()
	d := alg.NewDFS()

	d.Run(g, g.Init())

}

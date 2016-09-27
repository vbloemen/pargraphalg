package main

import (
	//"fmt"

	"github.com/vbloemen/pargraphalg/alg"
	"github.com/vbloemen/pargraphalg/graph"
)

func main() {

	a := graph.NewLine(5)
	b := graph.NewLine(5)
	g := graph.NewParallelComp(false, a, b)

	//g.PrintDOT()
	d := alg.NewDFS()
	d.Run(g, g.Init())

}

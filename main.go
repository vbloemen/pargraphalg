package main

import (
	"github.com/vbloemen/pargraphalg/alg"
	"github.com/vbloemen/pargraphalg/graph"
)

func main() {

	//	g := graph.TestGraph2()
	g := graph.NewComplete(5)
	d := alg.NewDFS()

	g.Print()
	d.Run(g, g.Init())

}

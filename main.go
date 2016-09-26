package main

import (
	"fmt"

	"github.com/vbloemen/pargraphalg/graph"
)

func main() {

	g := graph.TestGraph1()

	fmt.Println("Printing graph")
	g.Print()

	fmt.Println("\nPrinting graph in DOT format")
	g.PrintDOT()
}

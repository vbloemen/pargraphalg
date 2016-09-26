package main

import (
	"fmt"

	"github.com/vbloemen/pargraphalg/graph"
)

func main() {

	// 0 --> 1
	// \    /
	//  V  V
	//   2 

	testFrom := []int{0,2,3,3}
	testTo := []int{1,2,2}

	g := graph.Explicit{From: testFrom, To: testTo}

	fmt.Println("Init state: ", g.Init())
	for i := 0; i < g.NumStates(); i++ {
		fmt.Println("Successor(", i,"): ", g.Successors(i))
	}
}

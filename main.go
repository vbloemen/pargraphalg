package main

import (
	"fmt"
	"time"

	"github.com/vbloemen/pargraphalg/alg"
	"github.com/vbloemen/pargraphalg/graph"
)

func testSearch(g graph.Graph) {

	pbfs := alg.NewParBFS()
	fmt.Println("Starting Parallel BFS")
	start := time.Now()
	pbfs.Run(g, g.Init())
	elapsed := time.Since(start)
	fmt.Println("Done in", elapsed)

	pbfs = alg.NewParBFS()
	fmt.Println("Starting Parallel BFS Seq")
	start = time.Now()
	pbfs.RunSeq(g, g.Init())
	elapsed = time.Since(start)
	fmt.Println("Done in", elapsed)

	bfs := alg.NewBFS()
	fmt.Println("Starting BFS")
	start = time.Now()
	bfs.Run(g, g.Init())
	elapsed = time.Since(start)
	fmt.Println("Done in", elapsed)

	dfs := alg.NewDFS()
	fmt.Println("Starting DFS")
	start = time.Now()
	dfs.Run(g, g.Init())
	elapsed = time.Since(start)
	fmt.Println("Done in", elapsed)
}

func main() {

	a := graph.NewLoop(10)
	b := graph.NewLine(200)
	g := graph.NewParallelComp(false, a, b, b)

	//g.PrintDOT()
	testSearch(g)

	_ = b
	_ = g

	//fmt.Println("Number of states:",g.NumStates())
}

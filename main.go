package main

import (
	"fmt"
	"time"

	"github.com/vbloemen/pargraphalg/alg"
	"github.com/vbloemen/pargraphalg/graph"
)

func main() {
	g := graph.NewParallelComp(false, graph.NewLoop(10), graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	//g := graph.NewTree(25)

	fmt.Println("states:", g.NumStates())

	bfs := alg.NewBFS()
	fmt.Println("Starting Seq BFS")
	start := time.Now()
	bfs.Run(g, g.Init())
	fmt.Println("Done in", time.Since(start))

	dfs := alg.NewDFS()
	fmt.Println("Starting Seq DFS")
	start = time.Now()
	dfs.Run(g, g.Init())
	fmt.Println("Done in", time.Since(start))

	pbfs := alg.NewParBFS(4)
	fmt.Println("Starting Par BFS")
	start = time.Now()
	pbfs.Run(g, g.Init())
	fmt.Println("Done in", time.Since(start))
}

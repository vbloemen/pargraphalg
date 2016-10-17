package main

import (
	"fmt"
	"github.com/vbloemen/pargraphalg/alg"
	"github.com/vbloemen/pargraphalg/graph"
	_ "net/http/pprof"
	_ "runtime"
	"testing"
	"time"
)

const BENCH_TREE_SIZE = 25
const PAR_PROCS = 8

/*func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(4)
}*/

func TestBFSTree(t *testing.T) {
	g := graph.NewTree(BENCH_TREE_SIZE)
	bfs := alg.NewBFS()
	bfs.Run(g, g.Init())
}

func TestDFSTree(t *testing.T) {
	g := graph.NewTree(BENCH_TREE_SIZE)
	dfs := alg.NewDFS()
	dfs.Run(g, g.Init())
}

func TestParBFSTree(t *testing.T) {
	g := graph.NewTree(BENCH_TREE_SIZE)
	pbfs := alg.NewParBFSLB(PAR_PROCS)
	pbfs.Run(g, g.Init())
}

func TestDFSLi200Lo10(t *testing.T) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	dfs := alg.NewDFS()
	dfs.Run(g, g.Init())
}

// parallel BFS Tree

// parallel BFS Li200Lo10

func TestBFSLi200Lo10(t *testing.T) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	bfs := alg.NewBFS()

	fmt.Print("BFS       Li200Lo10 (", 1, "): ")
	start := time.Now()
	bfs.Run(g, g.Init())
	fmt.Println(time.Since(start))
}

func TestParBFSLi200Lo10(t *testing.T) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	pbfs := alg.NewParBFS(PAR_PROCS)

	fmt.Print("ParBFS    Li200Lo10 (", PAR_PROCS, "): ")
	start := time.Now()
	pbfs.Run(g, g.Init())
	fmt.Println(time.Since(start))
}

func TestParBFSLBLi200Lo10(t *testing.T) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	pbfs := alg.NewParBFSLB(PAR_PROCS)

	fmt.Print("ParBFSLB  Li200Lo10 (", PAR_PROCS, "): ")
	start := time.Now()
	pbfs.Run(g, g.Init())
	fmt.Println(time.Since(start))
}

func TestParBFSOSLi200Lo10(t *testing.T) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	pbfs := alg.NewParBFSOS(PAR_PROCS)

	fmt.Print("ParBFSOS  Li200Lo10 (", PAR_PROCS, "): ")
	start := time.Now()
	pbfs.Run(g, g.Init())
	fmt.Println(time.Since(start))
}

func TestParBFSVCLi200Lo10(t *testing.T) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	pbfs := alg.NewParBFSVC(PAR_PROCS)

	fmt.Print("ParBFSVC  Li200Lo10 (", PAR_PROCS, "): ")
	start := time.Now()
	pbfs.Run(g, g.Init())
	fmt.Println(time.Since(start))
}

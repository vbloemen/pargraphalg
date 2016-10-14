package main

import (
	"github.com/vbloemen/pargraphalg/alg"
	"github.com/vbloemen/pargraphalg/graph"
	_ "net/http/pprof"
	_ "runtime"
	"testing"
)

const BENCH_TREE_SIZE = 25
const PAR_PROCS = 4

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
	pbfs := alg.NewParBFSOS(PAR_PROCS)
	pbfs.Run(g, g.Init())
}

func TestBFSLi200Lo10(t *testing.T) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	bfs := alg.NewBFS()
	bfs.Run(g, g.Init())
}

func TestDFSLi200Lo10(t *testing.T) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	dfs := alg.NewDFS()
	dfs.Run(g, g.Init())
}

func TestParBFSLi200Lo10(t *testing.T) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200)) // todo change back
	pbfs := alg.NewParBFSOS(PAR_PROCS)
	pbfs.Run(g, g.Init())
}

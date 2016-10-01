package main

import (
	"testing"

	"github.com/vbloemen/pargraphalg/alg"
	"github.com/vbloemen/pargraphalg/graph"
)

const BENCH_TREE_SIZE = 25
const PAR_PROCS = 4

func BenchmarkBFSTree(b *testing.B) {
	g := graph.NewTree(BENCH_TREE_SIZE)
	bfs := alg.NewBFS()
	bfs.Run(g, g.Init())
}

func BenchmarkDFSTree(b *testing.B) {
	g := graph.NewTree(BENCH_TREE_SIZE)
	dfs := alg.NewDFS()
	dfs.Run(g, g.Init())
}

func BenchmarkParBFSTree(b *testing.B) {
	g := graph.NewTree(BENCH_TREE_SIZE)
	pbfs := alg.NewParBFS(PAR_PROCS)
	pbfs.Run(g, g.Init())
}

func BenchmarkBFSLi200Lo10(b *testing.B) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	bfs := alg.NewBFS()
	bfs.Run(g, g.Init())
}

func BenchmarkDFSLi200Lo10(b *testing.B) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	dfs := alg.NewDFS()
	dfs.Run(g, g.Init())
}

func BenchmarkParBFSLi200Lo10(b *testing.B) {
	g := graph.NewParallelComp(false, graph.NewLoop(10),
		graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	pbfs := alg.NewParBFS(PAR_PROCS)
	pbfs.Run(g, g.Init())
}

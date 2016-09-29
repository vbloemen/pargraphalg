package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/vbloemen/pargraphalg/alg"
	"github.com/vbloemen/pargraphalg/graph"
)

func testSearch(g graph.Graph) {

	pbfs2 := alg.NewParBFS2()
	fmt.Println("Starting Parallel BFS2")
	start := time.Now()
	pbfs2.Run(g, g.Init())
	elapsed := time.Since(start)
	fmt.Println("Done in", elapsed)

	return

	pbfs := alg.NewParBFS()
	fmt.Println("Starting Parallel BFS with Locking")
	start = time.Now()
	pbfs.RunWithLock(g, g.Init())
	elapsed = time.Since(start)
	fmt.Println("Done in", elapsed)

	pbfs = alg.NewParBFS()
	fmt.Println("Starting Parallel BFS")
	start = time.Now()
	pbfs.Run(g, g.Init())
	elapsed = time.Since(start)
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

var cpuprofile = flag.String("cpuprofile", "cpu.prof", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "mem.prof", "write memory profile to `file`")

func StartProfiling() {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
	}
}

func main() {
	flag.Parse()

	//a := graph.NewLoop(10)
	//b := graph.NewLine(200)
	//g := graph.NewParallelComp(false, a, a, b, b)

	//g.PrintDOT()
	//testSearch(g)

	g := graph.NewTree(22)
	fmt.Println("states:", g.NumStates())
	
	bfs := alg.NewBFS()
	fmt.Println("Starting BFS")
	start := time.Now()
	bfs.Run(g, g.Init())
	elapsed := time.Since(start)
	fmt.Println("Done in", elapsed)

	_ = g

	//fmt.Println("Number of states:",g.NumStates())

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}

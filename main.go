package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/vbloemen/pargraphalg/alg"
	"github.com/vbloemen/pargraphalg/graph"
)

// simple function that allows us to end the program at any time.
func interrupt() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Program exited with interrupt")
		os.Exit(1)
	}()
}

// convert an int to a string.
func toStr(val int, strlen int) string {
	return fmt.Sprintf("%2d", val)
}

func main() {
	interrupt()

	// parse the arguments
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Error: not a correct number of arguments")
		fmt.Println("Usage:")
		fmt.Println("      pargraphalg  GRAPH  ALG  THREADS")
		fmt.Println("e.g.  pargraphalg  tree25  bfslb  2")
		return
	}

	// graph
	var g graph.Graph
	switch args[0] {
	case "tree25":
		g = graph.NewTree(25)
	case "tree26":
		g = graph.NewTree(26)
	case "li200lo10":
		g = graph.NewParallelComp(false, graph.NewLoop(10),
			graph.NewLoop(10), graph.NewLine(200), graph.NewLine(200))
	default:
		fmt.Println("I don't know this graph!")
		fmt.Println("please select \"tree25\", \"tree26\" or \"li200lo10\"")
		return
	}

	// threads
	threads, err := strconv.Atoi(args[2])
	if err != nil || threads < 1 {
		fmt.Println("Unknown number of threads!")
		return
	} else if threads > runtime.NumCPU() {
		fmt.Println("Too many threads! Please select a maximum of ",
			runtime.NumCPU())
		return
	}

	// alg
	var a alg.Search
	switch args[1] {
	case "bfs":
		a = alg.NewBFS()
	case "dfs":
		a = alg.NewDFS()
	case "pbfs":
		a = alg.NewParBFS(threads)
	case "pbfsos":
		a = alg.NewParBFSOS(threads)
	case "pbfslb":
		a = alg.NewParBFSLB(threads)
	case "pbfsvc":
		a = alg.NewParBFSVC(threads)
	default:
		fmt.Println("I don't know this alg!")
		fmt.Println("please select \"bfs\", \"dfs\", \"pbfs\", \"pbfsos\", \"pbfslb\", \"pbfsvc\"")
		return
	}

	CSV_OUTPUT := true

	// run the algorithm on the graph
	//fmt.Printf("Starting %-10s with %-10s (%2d) : ", args[0], args[1], threads)
	if CSV_OUTPUT {
		fmt.Printf("%s,%s,%d,", args[0], args[1], threads) // CSV output
	} else {
		fmt.Printf("Starting %-10s with %-10s (%2d) : ", args[0], args[1], threads)
	}
	start := time.Now()

	a.Run(g, g.Init())

	fmt.Println(time.Since(start).Seconds())
}

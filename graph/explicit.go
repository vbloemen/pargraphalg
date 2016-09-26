package graph

import (
	"fmt"
)

// An explicitly given graph implementation, using the CSR style of two arrays
// for representing the successors.
type Explicit struct {
	Graph // implementing the Graph interface
	From  []int
	To    []int
}

// Returns the initial state of the graph, 0.
func (g Explicit) Init() int {
	return 0
}

// Returns a slice of the successors for a state.
func (g Explicit) Successors(state int) []int {
	return g.To[g.From[state]:g.From[state+1]]
}

// Returns the number of state the graph.
func (g Explicit) NumStates() int {
	return len(g.From) - 1
}

// Prints the graph to standard output.
func (g Explicit) Print() {
	for i := 0; i < g.NumStates(); i++ {
		fmt.Println("Successors", i, ":", g.Successors(i))
	}
}

// Prints the graph to standard output in DOT format.
func (g Explicit) PrintDOT() {
	fmt.Println("digraph graph {")
	for i := 0; i < g.NumStates(); i++ {
		for _, suc := range g.Successors(i) {
			fmt.Printf("  %d -> %d\n", i, suc)
		}
	}
	fmt.Println("}")
}

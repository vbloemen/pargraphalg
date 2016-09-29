package graph

import (
	"fmt"
)

// A tree graph of a specified depth
type Tree struct {
	Graph     // implementing the Graph interface
	n     int // the number of states in the line
}

// Constructor for the Tree graph type.
func NewTree(depth int) *Tree {
	return &Tree{n: 1<<uint(depth) - 1}
}

// Returns the initial state of the graph, 0.
func (g Tree) Init() int {
	return 0
}

// Returns a slice of the successors for a state.
func (g Tree) Successors(state int) []int {
	suc := state*2 + 1
	if suc >= g.NumStates() {
		return []int{}
	}
	return []int{suc, suc + 1}
}

// Returns the number of state the graph.
func (g Tree) NumStates() int {
	return g.n
}

// Prints the graph to standard output.
func (g Tree) Print() {
	for i := 0; i < g.NumStates(); i++ {
		fmt.Println("Successors", i, ":", g.Successors(i))
	}
}

// Prints the graph to standard output in DOT format.
func (g Tree) PrintDOT() {
	fmt.Printf("// number of states:      %d\n", g.NumStates())
	fmt.Println("digraph g {")
	for i := 0; i < g.NumStates(); i++ {
		fmt.Printf("  %d [shape=circle];\n", i)
		for _, suc := range g.Successors(i) {
			fmt.Printf("  %d -> %d;\n", i, suc)
		}
	}
	fmt.Println("}")
}

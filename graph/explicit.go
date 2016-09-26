package graph

import (
	"fmt"
)

// An explicitly given graph implementation, using the CSR style of two arrays
// for representing the successors.
type Explicit struct {
	Graph       // implementing the Graph interface
	From  []int /* contains starting indices for the successors.
	   NB: end with an extra entry with value len(To). */
	To []int
}

// Constructor for the Explicit type.
func NewExplicit(from []int, to []int) *Explicit {
	return &Explicit{From: from, To: to}
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
	fmt.Printf("// number of states:      %d\n", g.NumStates())
	fmt.Printf("// number of transitions: %d\n", len(g.To))
	fmt.Println("digraph g {")
	for i := 0; i < g.NumStates(); i++ {
		fmt.Printf("  %d [shape=circle];\n", i)
		for _, suc := range g.Successors(i) {
			fmt.Printf("  %d -> %d;\n", i, suc)
		}
	}
	fmt.Println("}")
}

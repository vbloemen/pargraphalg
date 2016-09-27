package graph

import (
	"fmt"
)

// A complete Kn graph for a specific n
type Complete struct {
	Graph     // implementing the Graph interface
	n     int // the number of states
}

// Constructor for the Complete graph type.
func NewComplete(nstates int) *Complete {
	return &Complete{n: nstates}
}

// Returns the initial state of the graph, 0.
func (g Complete) Init() int {
	return 0
}

// Returns a slice of the successors for a state. This consists of all other
// states, so excluding the current one.
func (g Complete) Successors(state int) []int {
	sucs := make([]int, g.NumStates()-1)
	for i := 0; i < state; i++ {
		sucs[i] = i
	}
	for i := state + 1; i < g.NumStates(); i++ {
		sucs[i-1] = i
	}
	return sucs
}

// Returns the number of state the graph.
func (g Complete) NumStates() int {
	return g.n
}

// Prints the graph to standard output.
func (g Complete) Print() {
	for i := 0; i < g.NumStates(); i++ {
		fmt.Println("Successors", i, ":", g.Successors(i))
	}
}

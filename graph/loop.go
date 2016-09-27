package graph

import (
	"fmt"
)

// A loop graph of a specified length
type Loop struct {
	Graph     // implementing the Graph interface
	n     int // the number of states in the loop
}

// Constructor for the Loop graph type.
func NewLoop(nstates int) *Loop {
	return &Loop{n: nstates}
}

// Returns the initial state of the graph, 0.
func (g Loop) Init() int {
	return 0
}

// Returns a slice of the successors for a state.
func (g Loop) Successors(state int) []int {
	return []int{(state + 1) % g.NumStates()}
}

// Returns the number of state the graph.
func (g Loop) NumStates() int {
	return g.n
}

// Prints the graph to standard output.
func (g Loop) Print() {
	for i := 0; i < g.NumStates(); i++ {
		fmt.Println("Successors", i, ":", g.Successors(i))
	}
}

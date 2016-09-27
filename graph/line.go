package graph

import (
	"fmt"
)

// A line graph of a specified length
type Line struct {
	Graph     // implementing the Graph interface
	n     int // the number of states in the line
}

// Constructor for the Line graph type.
func NewLine(nstates int) *Line {
	return &Line{n: nstates}
}

// Returns the initial state of the graph, 0.
func (g Line) Init() int {
	return 0
}

// Returns a slice of the successors for a state.
func (g Line) Successors(state int) []int {
	if state == g.NumStates()-1 {
		return []int{}
	}
	return []int{state + 1}
}

// Returns the number of state the graph.
func (g Line) NumStates() int {
	return g.n
}

// Prints the graph to standard output.
func (g Line) Print() {
	for i := 0; i < g.NumStates(); i++ {
		fmt.Println("Successors", i, ":", g.Successors(i))
	}
}

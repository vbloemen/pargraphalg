package graph

// An explicitly given graph implementation, using the CSR style of two arrays 
// for representing the successors.
type Explicit struct {
	Graph // implementing the Graph interface
	From []int
	To []int
}


// Returns the initial state of the graph, 0.
func (g Explicit) Init() int {
    return 0
}


// Returns a slice of the successors for a state.
func (g Explicit) Successors(state int) []int {
    return g.To[ g.From[state] : g.From[state+1] ]
}


// Returns the number of state the graph.
func (g Explicit) NumStates() int {
    return len(g.From) - 1
}

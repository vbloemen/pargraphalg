package graph

import (
	"fmt"
)

// A parallel composition of multiple given subgraphs. Note that the state
// numbers are composed from the ones in the subgraphs, which assumes that these
// are labeled from 0 to n.
type ParallelComp struct {
	Graph         // implementing the Graph interface
	gs    []Graph // the subgraphs
	ngs   int     // the number of subgraphs
	nsgs  []int   // the number of states per subgraph
	fgs   []int   // the applied factor per graph for computing unique
	// state numbers
	n         int // the total number of states
	allowSync bool
}

// Constructor for the ParallelComp graph type. It computes the number of states
// for each subgraph, to be used for uniquely identifying the states. the
// allowsync bool determines whether a successor computation may perform the
// successors of multiple graphs simultaneously.
func NewParallelComp(allowsync bool, graphs ...Graph) *ParallelComp {
	nsgraphs := make([]int, len(graphs))
	fgraphs := make([]int, len(graphs))
	totalStates := 1
	for i, gi := range graphs {
		nsgraphs[i] = gi.NumStates()
		fgraphs[i] = totalStates
		totalStates *= nsgraphs[i]
	}

	return &ParallelComp{gs: graphs, ngs: len(graphs), nsgs: nsgraphs,
		fgs: fgraphs, n: totalStates, allowSync: allowsync}
}

// Returns the initial state as a .
func (g ParallelComp) Init() int {
	state := 0
	for i, gi := range g.gs {
		state += gi.Init() * g.fgs[i]
	}
	return state
}

// Combines the sets of successors from each subgraph in every possible
// combination. Note that this function also returns one extra successor, for
// which none of the subgraphs 'take a step'.
func (g ParallelComp) combineSuccessorsSync(sucs [][]int, csucs []int,
	ci *int, gi, cursuc int) {

	// end of recursion: store successor in csucs
	if gi >= g.ngs {
		csucs[*ci] = cursuc
		*ci++
		return
	}

	// iterate over all successors in this subgraph
	for _, sj := range sucs[gi] {
		// when taking successor, don't just add the successor of the subgraph,
		// also remove the old value for that subgraph
		g.combineSuccessorsSync(sucs, csucs, ci, gi+1,
			cursuc+(sj-g.StateNumToArr(cursuc)[gi])*g.fgs[gi]) // take
	}
	g.combineSuccessorsSync(sucs, csucs, ci, gi+1, cursuc) // don't take
}

// Combines the sets of successors from each subgraph in every possible
// combination. Note that this function also returns one extra successor, for
// which none of the subgraphs 'take a step'.
func (g ParallelComp) combineSuccessors(sucs [][]int, csucs []int,
	ci *int, gi, state int) {

	// end of recursion: store successor in csucs
	if gi >= g.ngs {
		return
	}

	// iterate over all successors in this subgraph
	for _, sj := range sucs[gi] {
		// when taking successor, don't just add the successor of the subgraph,
		// also remove the old value for that subgraph
		csucs[*ci] = state + (sj-g.StateNumToArr(state)[gi])*g.fgs[gi]
		*ci++
	}
	g.combineSuccessors(sucs, csucs, ci, gi+1, state) // don't take
}

// Returns a slice of the successors for a state, by first computing all
// possible successors of each component and combining these for every
// combination.
func (g ParallelComp) Successors(state int) []int {
	substates := g.StateNumToArr(state)
	// compute the successors for each subgraph
	sucs := make([][]int, g.ngs)
	totalsizeSync := 1
	totalsize := 1

	for i, si := range substates {
		sucs[i] = g.gs[i].Successors(si)
		totalsizeSync *= (len(sucs[i]) + 1)
		totalsize += len(sucs[i])
	}
	if g.allowSync {
		csucs := make([]int, totalsizeSync)
		ci := 0
		g.combineSuccessorsSync(sucs, csucs, &ci, 0, state)
		return csucs[:totalsizeSync-1]
	}

	csucs := make([]int, totalsize)
	ci := 0
	g.combineSuccessors(sucs, csucs, &ci, 0, state)
	return csucs[:totalsize-1]
}

// Returns the number of state the graph if it is set.
func (g ParallelComp) NumStates() int {
	return g.n
}

// Auxiliary function to identify and return the state indices of each subgraph.
func (g ParallelComp) StateNumToArr(state int) []int {
	arr := make([]int, g.ngs)
	s := state
	for i := g.ngs - 1; i >= 0; i-- {
		if s >= g.fgs[i] {
			arr[i] += s / g.fgs[i]
			s = s % g.fgs[i]
		}
	}
	return arr
}

// Auxiliary function to compute the state number from the indices of each
// subgraph. Assumes that len(states) == len(g.gs).
func (g ParallelComp) StateArrToNum(states []int) int {
	state := 0
	for i, si := range states {
		state += si * g.fgs[i]
	}
	return state
}

// Prints the graph to standard output in DOT format.
func (g ParallelComp) PrintDOT() {
	fmt.Printf("// number of states:      %d\n", g.NumStates())
	fmt.Println("digraph g {")
	for i := 0; i < g.NumStates(); i++ {
		fmt.Printf("  \"%v\" [shape=circle];\n", g.StateNumToArr(i))
		for _, suc := range g.Successors(i) {
			fmt.Printf("  \"%v\" -> \"%v\";\n", g.StateNumToArr(i),
				g.StateNumToArr(suc))
		}
	}
	fmt.Println("}")
}

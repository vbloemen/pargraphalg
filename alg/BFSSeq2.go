package alg

import (
	"fmt"

	"github.com/vbloemen/pargraphalg/graph"
)

// Data type for BFSSeq2.
type BFSSeq2 struct {
	Search        // implementing the Search interface
	V      []bool // visited set
	C      []int  // queue array
}

// Constructor for the BFS type.
func NewBFSSeq2() *BFSSeq2 {
	C := make([]int, 1e8)
	V := make([]bool, 1e8)
	return &BFSSeq2{C: C, V: V}
}

// best sequential version of BFSSeq2,
func (b *BFSSeq2) Run(g graph.Graph, from int) {
	// init search setup
	b.C[0] = from
	b.V[from] = true
	Ci := 0 // queue index
	Cn := 1 // queue length

	for Ci < Cn {
		sucs := g.Successors(b.C[Ci])
		for _, si := range sucs {
			if !b.V[si] {
				b.V[si] = true
				b.C[Cn] = si // add the state to the queue
				Cn++
			}
		}
		Ci++
	}
	fmt.Println("State count and actual", Cn, g.NumStates())
}

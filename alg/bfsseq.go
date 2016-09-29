package alg

import (
	"fmt"

	"github.com/vbloemen/pargraphalg/graph"
)

// Data type for BFSSeq.
type BFSSeq struct {
	Search  // implementing the Search interface
	V                 []bool                     // visited set
	C                 chan int                   // current queue channel
}

// Constructor for the BFS type.
func NewBFSSeq() *BFSSeq {
	C := make(chan int, 1e7)
	V := make([]bool, 1e7)
	return &BFSSeq{C: C, V: V}
}

// best sequential version of BFSSeq, 
func (b *BFSSeq) Run(g graph.Graph, from int) {
	// init search setup
	b.C <- from
	stateCount := 0

	// check and update visited states with multiple goroutines
	for {
		select {
		case state := <-b.C:
			sucs := g.Successors(state)
			for _, si := range sucs {
				if !b.V[si] {
					b.V[si] = true
					b.C <- si // add the state to the queue
				}
			}
			stateCount ++
		default:
			fmt.Println("State count and actual",stateCount,g.NumStates())
			return
		}
	}
}

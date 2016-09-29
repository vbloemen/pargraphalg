package alg

import (
	"fmt"

	"github.com/vbloemen/pargraphalg/graph"
)

var _ = fmt.Printf // For debugging; delete when done.

// NB: for now we don't need a visited set

// Data type for ParBFS5. This one doesn't use queues.
type ParBFS5 struct {
	Search                                       // implementing the Search interface
	V                 []bool                     // visited set
	C                 chan int                   // current queue channel
	VC                [maxVisitHandlers]chan int // visit channels
	spawnedComputes   int
	completedComputes int
}

// Constructor for the BFS type.
func NewParBFS5() *ParBFS5 {
	C := make(chan int, 1e7)
	V := make([]bool, 1e7)

	var vc [maxVisitHandlers]chan int
	for i := range vc {
		vc[i] = make(chan int)
	}
	return &ParBFS5{V: V, C: C, VC: vc, spawnedComputes: 0, completedComputes: 0}
}

func (b *ParBFS5) ComputeSuccessors(g graph.Graph, state int, doneComputeChannel chan bool) {

	sucs := g.Successors(state)
	for _, si := range sucs {
		b.C <- si // add the state to the queue
	}
	doneComputeChannel <- true
}

// Performs a parallel BFS by using two queues and swapping the queues in every
// iteration. Implemented with a sync WaitGroup
func (b *ParBFS5) Run(g graph.Graph, from int) {
	// init search setup
	b.C <- from
	b.V[from] = true
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
			stateCount++
		default:
			fmt.Println("State count and actual", stateCount, g.NumStates())
			return
		}
	}
}

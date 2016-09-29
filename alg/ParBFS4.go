package alg

import (
	"fmt"

	"github.com/vbloemen/pargraphalg/graph"
)

var _ = fmt.Printf // For debugging; delete when done.
const maxVisitHandlers = 100

// Data type for ParBFS4. This one doesn't use queues.
type ParBFS4 struct {
	Search                                       // implementing the Search interface
	V                 []bool                     // visited set
	C                 chan int                   // current queue channel
	VC                [maxVisitHandlers]chan int // visit channels
	spawnedComputes   int
	completedComputes int
}

// Constructor for the BFS type.
func NewParBFS4() *ParBFS4 {
	C := make(chan int, 1e7)
	V := make([]bool, 1e7)

	var vc [maxVisitHandlers]chan int
	for i := range vc {
		vc[i] = make(chan int)
	}
	return &ParBFS4{V: V, C: C, VC: vc, spawnedComputes: 0, completedComputes: 0}
}

func (b *ParBFS4) ComputeSuccessors(g graph.Graph, state int, doneComputeChannel chan bool) {

	sucs := g.Successors(state)
	for _, si := range sucs {
		b.VC[si%maxVisitHandlers] <- si
	}
	doneComputeChannel <- true
}

func (b *ParBFS4) handleVisited(i int) {
	for {
		state := <-b.VC[i]
		if state == -1 {
			return
		}
		if !b.V[state] {
			b.V[state] = true
			b.C <- state // add the state to the queue
		}
	}
}

// Performs a parallel BFS by using two queues and swapping the queues in every
// iteration. Implemented with a sync WaitGroup
func (b *ParBFS4) Run(g graph.Graph, from int) {

	doneComputeChannel := make(chan bool, 1e8)

	// init search setup
	b.C <- from
	b.V[from] = true
	stateCount := 0
	spawned := 0
	doneComputing := 0

	for i := 0; i < maxVisitHandlers; i++ {
		go b.handleVisited(i)
	}
	// check and update visited states with multiple goroutines
	for {
		select {
		case state := <-b.C:
			go b.ComputeSuccessors(g, state, doneComputeChannel)
			spawned++
			stateCount++
		default:
			if spawned > doneComputing {
				<-doneComputeChannel
				doneComputing++
			} else {
				fmt.Println("spawned", spawned, "done", doneComputing)
				// buffer empty, thus we have ended
				fmt.Println("State count and actual:", stateCount, g.NumStates())

				// close the visiter handlers
				for i := 0; i < maxVisitHandlers; i++ {
					b.VC[i] <- -1
				}

				return
			}
		}
	}

}

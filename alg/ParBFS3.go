package alg

import (
	"fmt"

	"github.com/vbloemen/pargraphalg/graph"
)

var _ = fmt.Printf // For debugging; delete when done.
const maxVisitHandlers = 100

// Data type for ParBFS3.
type ParBFS3 struct {
	Search        // implementing the Search interface
	V      []bool // visited set
	C      []int  // current queue
	Cn     int  // current queue index
	N      []int  // next queue
	Nn     int  // next queue index

	// all channels
	VC [maxVisitHandlers]chan int
}

// Constructor for the BFS type.
func NewParBFS3() *ParBFS3 {
	C := make([]int, 1e7)
	N := make([]int, 1e7)
	V := make([]bool, 1e7)

	var visitedChan [maxVisitHandlers]chan int
	for i := range visitedChan {
		visitedChan[i] = make(chan int)
	}
	return &ParBFS3{V: V, C: C, N: N, Cn: 0, Nn: 0, VC: visitedChan}
}

func (b *ParBFS3) handleState(g graph.Graph, state int, doneChan chan bool) {

	sucs := g.Successors(state)
	for _, si := range sucs {
		b.VC[si%maxVisitHandlers] <- si
	}
	doneChan <- true
}

func (b *ParBFS3) handleNewStates(newStateChan chan int, flushChan chan bool) {
	for {
		state := <-newStateChan
		if state == -1 {
			flushChan <- true
			continue
		}
		b.Nn++
		b.N[b.Nn-1] = state
	}
}

func (b *ParBFS3) handleVisited(i int, newStateChan chan int, flushChan chan bool) {
	for {
		state := <-b.VC[i]
		if state == -1 {
			flushChan <- true
			continue
		}
		if !b.V[state] {
			b.V[state] = true
			newStateChan <- state // add the state to the queue
		}
	}
}

// Performs a parallel BFS by using two queues and swapping the queues in every
// iteration. Implemented with a sync WaitGroup
func (b *ParBFS3) Run(g graph.Graph, from int) {
	// init setup
	b.C[0] = from
	b.Cn = 1
	b.V[from] = true
	doneChan := make(chan bool)
	flushChan := make(chan bool)
	newStateChan := make(chan int)

	// sequentially add new states to the queue
	go b.handleNewStates(newStateChan, flushChan)

	// check and update visited states with multiple goroutines
	for i := 0; i < maxVisitHandlers; i++ {
		go b.handleVisited(i, newStateChan, flushChan)
	}

	for b.Cn > 0 {

		// do the successor call for every state in parallel
		numRoutines := b.Cn
		for i := 0; i < b.Cn; i++ {
			go b.handleState(g, b.C[i], doneChan)
		}

		// wait until all successors have been put on the VC[] channels
		for i := 0; i < numRoutines; i++ {
			<-doneChan
		}

		// wait until the VC[] channels are done
		go func() {
			for i := 0; i< maxVisitHandlers; i++ {
				b.VC[i] <- -1
			}
		}()
		// get notified that the VC channels are done
		for i := 0; i< maxVisitHandlers; i++ {
			<-flushChan 
		}


		// TODO: Keep track of all solutions (probably interesting for a paper)
		// TODO: Multiple queues for processing?
		// TODO: Check how LTSmin does it



		// flush the newStateChan
		go func() {
			newStateChan <- -1
		}()
		// and get notified that it is flushed
		<-flushChan


		// swapping queues can be done in parallel, but hopefully we can
		// just do pointer swapping
		b.N, b.C = b.C, b.N
		b.Cn = b.Nn
		b.Nn = 0
	}

	// TODO: cleanup using doneHandler

}

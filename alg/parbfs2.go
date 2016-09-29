package alg

import (
	"fmt"
	//"sync"
	//"sync/atomic"
	//"runtime"

	"github.com/vbloemen/pargraphalg/graph"
)

var _ = fmt.Printf // For debugging; delete when done.
const maxVisitHandlers2 = 1000

// Data type for ParBFS2.
type ParBFS2 struct {
	Search        // implementing the Search interface
	V      []bool // visited set
	C      []int  // current queue
	Cn     int32  // current queue index
	N      []int  // next queue
	Nn     int32  // next queue index
}

// Constructor for the BFS type.
func NewParBFS2() *ParBFS2 {
	C := make([]int, 1e7)
	N := make([]int, 1e7)
	V := make([]bool, 1e7)
	return &ParBFS2{V: V, C: C, N: N, Cn: 0, Nn: 0}
}

// returns whether state has already been visited, otherwise it returns false
// and adds state to the visited set
func (b *ParBFS2) FindOrPut(state int) bool {
	if b.V[state] {
		return true
	}
	b.V[state] = true
	return false
}

func (b *ParBFS2) handleState(g graph.Graph, state int, visitedChan *[maxVisitHandlers2]chan int, doneChan chan bool) {

	sucs := g.Successors(state)
	for _, si := range sucs {
		visitedChan[si%maxVisitHandlers2] <- si
	}
	doneChan <- true
}

func (b *ParBFS2) handleNewStates(newStateChan chan int, doneHandler chan bool) {
	for {
		select {
		case state := <-newStateChan:
			b.Nn++
			b.N[b.Nn-1] = state
		case <-doneHandler:
			return
		}
	}
}

func (b *ParBFS2) handleVisited(i int, visitedChan chan int, newStateChan chan int, doneHandler chan bool) {
	for {
		select {
		case state := <-visitedChan:
			if !b.V[state] {
				b.V[state] = true
				newStateChan <- state
			}
		case <-doneHandler:
			return
		}
	}
}

// Performs a parallel BFS by using two queues and swapping the queues in every
// iteration. Implemented with a sync WaitGroup
func (b *ParBFS2) Run(g graph.Graph, from int) {
	// init setup
	b.C[0] = from
	b.Cn = 1
	b.FindOrPut(from)
	doneChan := make(chan bool)
	newStateChan := make(chan int)
	doneHandler := make(chan bool)

	var visitedChan [maxVisitHandlers2]chan int
	for i := range visitedChan {
		visitedChan[i] = make(chan int)
	}

	go b.handleNewStates(newStateChan, doneHandler)

	for i := 0; i < maxVisitHandlers2; i++ {
		go b.handleVisited(i, visitedChan[i], newStateChan, doneHandler)
	}

	for b.Cn > 0 {

		// do the successor call for every state in parallel
		numRoutines := b.Cn
		var i int32
		for i = 0; i < b.Cn; i++ {
			go b.handleState(g, b.C[i], &visitedChan, doneChan)
		}

		// wait until all processes have finished
		for i = 0; i < numRoutines; i++ {
			<-doneChan
		}

		// swapping queues can be done in parallel, but hopefully we can
		// just do pointer swapping
		b.N, b.C = b.C, b.N
		b.Cn = b.Nn
		b.Nn = 0
	}

	// TODO: cleanup using doneHandler

}

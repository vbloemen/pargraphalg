package alg

import (
	"fmt"
	"sync"
	//"runtime"

	"github.com/vbloemen/pargraphalg/graph"

	"github.com/karalabe/cookiejar/collections/queue"
)

var _ = fmt.Printf // For debugging; delete when done.

var mu = &sync.Mutex{}

// Data type for ParBFS.
type ParBFS struct {
	Search  // implementing the Search interface
	visited map[int]bool
	q1      *queue.Queue
	q2      *queue.Queue
	mutex   sync.Mutex
}

// Constructor for the BFS type.
func NewParBFS() *ParBFS {
	q1 := queue.New()
	q2 := queue.New()
	return &ParBFS{visited: make(map[int]bool), q1: q1, q2: q2}
}

func (b ParBFS) handleState(g graph.Graph, state int, sucChan chan int,
	doneChan chan bool) {
	sucs := g.Successors(state)
	for _, si := range sucs {
		//fmt.Println("found suc", si)
		sucChan <- si
	}
	doneChan <- true
}

func (b ParBFS) handleStateLock(g graph.Graph, state int, doneChan chan bool) {

	sucs := g.Successors(state)
	for _, si := range sucs {
		// avoid concurrent read and write
		mu.Lock()

		if !b.visited[si] {
			b.visited[si] = true
			b.q2.Push(si)
		}

		mu.Unlock()
		//runtime.Gosched()
	}
	doneChan <- true
}

// Performs a parallel BFS by using two queues and swapping the queues in every
// iteration. Implemented with a sync WaitGroup
func (b ParBFS) RunWithLock(g graph.Graph, from int) {
	// init setup
	b.q1.Push(from)
	b.visited[from] = true
	//var wg sync.WaitGroup
	doneChan := make(chan bool)

	for !b.q1.Empty() {

		// do the successor call for every state in parallel
		numRoutines := 0
		for !b.q1.Empty() {
			state := b.q1.Pop().(int)
			numRoutines++
			go b.handleStateLock(g, state, doneChan)
		}

		// wait until all processes have finished
	Loop:
		for {
			select {
			case <-doneChan:
				numRoutines--
				if numRoutines == 0 {
					break Loop
				}

			}
		}

		b.q1, b.q2 = b.q2, b.q1

	}

}

// Performs a parallel BFS by using two queues and swapping the queues in every
// iteration.
func (b ParBFS) Run(g graph.Graph, from int) {
	b.q1.Push(from)
	b.visited[from] = true
	sucChan := make(chan int)
	doneChan := make(chan bool)

	for !b.q1.Empty() {

		// do the successor call for every state in parallel
		numRoutines := 0
		for !b.q1.Empty() {
			state := b.q1.Pop().(int)
			numRoutines++
			go b.handleState(g, state, sucChan, doneChan)
		}

		// handle visited and push sequentially, for now
	Loop:
		for {
			select {
			case si := <-sucChan:
				if !b.visited[si] {
					b.visited[si] = true
					b.q2.Push(si)
				}
			case <-doneChan:
				//fmt.Println("done!", msg, numRoutines)
				numRoutines--
				if numRoutines == 0 {
					break Loop
				}

			}
		}

		b.q1, b.q2 = b.q2, b.q1

	}

}

// Performs a parallel BFS in a sequential setting by using two queues and
// swapping the queues in every iteration.
func (b ParBFS) RunSeq(g graph.Graph, from int) {
	b.q1.Push(from)
	b.visited[from] = true

	for !b.q1.Empty() {
		for !b.q1.Empty() {
			state := b.q1.Pop().(int)
			sucs := g.Successors(state)
			for _, si := range sucs {
				if !b.visited[si] {
					b.visited[si] = true
					b.q2.Push(si)
				}
			}
		}

		b.q1, b.q2 = b.q2, b.q1

	}

}

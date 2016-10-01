package alg

import (
	"fmt"
	//"runtime"
	"sync"

	"github.com/vbloemen/pargraphalg/graph"
)

// Data type for ParBFS5.
type ParBFS5 struct {
	Search         // implementing the Search interface
	V      []bool  // visited set
	C      []int64 // current array
	N      []int64 // next array
	Ci     int64
	Cn     int64
	Ni     int64
	Nn     int64
	nProcs int64
	mu     *sync.Mutex
}

// Constructor for the BFS type.
func NewParBFS5() *ParBFS5 {
	C := make([]int64, 1e8)
	N := make([]int64, 1e8)
	V := make([]bool, 1e8)
	return &ParBFS5{C: C, N: N, V: V, Ci: 0, Cn: 0, Ni: 0, Nn: 0, nProcs: 0,
		mu: &sync.Mutex{}}
}

func (b *ParBFS5) proc(g graph.Graph, from int64, to int64, done chan bool) {

	for i := from; i < to; i++ {
		sucs := g.Successors(int(b.C[i]))
		var si int64
		for _, ssi := range sucs {
			si = int64(ssi)

			b.mu.Lock()
			if !b.V[si] {
				b.V[si] = true
				b.N[b.Nn] = si // add the state to the queue
				b.Nn++
			}
			b.mu.Unlock()
		}
	}
	done <- true
}

// Spawn X processes that all process the current layer in parallel, by
// distributing the work evenly: [0..Cn/X, Cn/X+1.., ..Cn].
// Once they're done, it reports this on the channel.
func (b *ParBFS5) Run(g graph.Graph, from int) {
	// init search setup
	b.C[0] = int64(from)
	b.V[from] = true
	b.Ci = 0 // current queue index
	b.Cn = 1 // current queue length
	b.Ni = 0 // next queue index
	b.Nn = 0 // next queue length
	var stateCount int64 = 0
	procs := 1 // somehow this is faster than 8 procs

	done := make(chan bool, procs)

	for b.Cn > 0 {
		step := int(b.Cn)/procs
		for p := 0; p < procs-1; p++ {
			//fmt.Println("Starting",int64(p*step),"to",int64((p+1)*step+1),"max:",b.Cn)
			go b.proc(g, int64(p*step), int64((p+1)*step), done)
		}
		//fmt.Println("Starting",int64((procs-1)*step),"to",b.Cn,"max:",b.Cn)
		go b.proc(g, int64((procs-1)*step), b.Cn, done)

		// wait for all procs to finish
		for p := 0; p < procs; p++ {
			<-done
		}

		//fmt.Println("Finished iteration of", b.Cn, "states")

		stateCount += b.Cn
		b.C, b.N = b.N, b.C
		b.Ci = 0
		b.Ni = 0
		b.Cn = b.Nn
		b.Nn = 0
	}

	fmt.Println("State count and actual", stateCount, g.NumStates(), procs)
	if stateCount != int64(g.NumStates()) {
		panic("Wrong number of states!")
	}
}

package alg

import (
	//"fmt"
	//"runtime"
	"sync"
	"sync/atomic"

	"github.com/vbloemen/pargraphalg/graph"
)

// Data type for ParBFS.
type ParBFS struct {
	Search         // implementing the Search interface
	V      []int64 // visited set
	C      []int64 // current queue
	N      []int64 // next queue
	Ci     int64
	Cn     int64
	Ni     int64
	Nn     int64
	mu     *sync.Mutex
	procs  int // number of goroutines to spawn
}

// Constructor for the BFS type.
func NewParBFS(procs int) *ParBFS {
	C := make([]int64, 1e8)
	N := make([]int64, 1e8)
	V := make([]int64, 1e8)
	return &ParBFS{C: C, N: N, V: V, Ci: 0, Cn: 0, Ni: 0, Nn: 0, procs: procs,
		mu: &sync.Mutex{}}
}

func (b *ParBFS) proc(g graph.Graph, from int64, to int64, done chan bool) {
	//runtime.LockOSThread()
	for i := from; i < to; i++ {
		sucs := g.Successors(int(b.C[i]))
		var si int64
		for _, ssi := range sucs {
			si = int64(ssi)

			if atomic.CompareAndSwapInt64(&b.V[si], 0, 1) {
				newN := atomic.AddInt64(&b.Nn, 1)
				b.N[newN-1] = si // add the state to the queue
			}

			//          // mutex lock approach
			//			b.mu.Lock()
			//			if b.V[si] == 0 {
			//				b.V[si] = 1
			//				b.N[b.Nn] = si // add the state to the queue
			//				b.Nn++
			//			}
			//			b.mu.Unlock()
		}
	}
	done <- true
}

// Spawn X processes that all process the current layer in parallel, by
// distributing the work evenly: [0..Cn/X) [Cn/X..2*Cn/X) .. [(X-1)Cn/X..Cn).
// Once they're done, it reports this on the 'done' channel. The main proc
// will wait for everything to finish, swap the current and next queues and
// start again.
func (b *ParBFS) Run(g graph.Graph, from int) {
	// init search setup
	b.C[0] = int64(from)
	b.V[from] = 1
	b.Ci = 0 // current queue index
	b.Cn = 1 // current queue length
	b.Ni = 0 // next queue index
	b.Nn = 0 // next queue length
	var stateCount int64 = 0

	done := make(chan bool, b.procs)

	for b.Cn > 0 {
		step := int(b.Cn) / b.procs
		for p := 0; p < b.procs-1; p++ {
			//fmt.Println("Starting",int64(p*step),"to",int64((p+1)*step+1),"max:",b.Cn)
			go b.proc(g, int64(p*step), int64((p+1)*step), done)
		}
		//fmt.Println("Starting",int64((b.procs-1)*step),"to",b.Cn,"max:",b.Cn)
		go b.proc(g, int64((b.procs-1)*step), b.Cn, done)

		// wait for all b.procs to finish
		for p := 0; p < b.procs; p++ {
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

	//fmt.Println("State count and actual", stateCount, g.NumStates(),
	//	"for", b.procs, "procs")
	if stateCount != int64(g.NumStates()) {
		panic("Wrong number of states!")
	}
}

package alg

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/vbloemen/pargraphalg/graph"
)

// Data type for ParBFSLB.
type ParBFSLB struct {
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
func NewParBFSLB(procs int) *ParBFSLB {
	C := make([]int64, 1e8)
	N := make([]int64, 1e8)
	V := make([]int64, 1e8)
	return &ParBFSLB{C: C, N: N, V: V, Ci: 0, Cn: 0, Ni: 0, Nn: 0, procs: procs,
		mu: &sync.Mutex{}}
}

// rangeC is the input channel
// doneC is the output channel
func (b *ParBFSLB) proc(g graph.Graph, doneC chan bool, rangeC chan int64) {
	runtime.LockOSThread()
	var buffer_size int64 = 10
	buffer := make([]int64, buffer_size)
	var bn int64 = 0
	var bi int64 = 0

	for {
		from := <-rangeC
		to := <-rangeC
		for i := from; i < to; i++ {
			sucs := g.Successors(int(b.C[i]))
			var si int64
			for _, ssi := range sucs {
				si = int64(ssi)

				// locally write to buffer, if its full, add to queue
				if atomic.CompareAndSwapInt64(&b.V[si], 0, 1) {
					if bn == buffer_size {
						newN := atomic.AddInt64(&b.Nn, bn) - bn
						for bi = 0; bi < bn; bi++ {
							b.N[newN+int64(bi)] = buffer[bi]
						}
						bn = 0
					} 
					buffer[bn] = si
					bn++
				}

				//if atomic.CompareAndSwapInt64(&b.V[si], 0, 1) {
				//	newN := atomic.AddInt64(&b.Nn, 1)
				//	b.N[newN-1] = si // add the state to the queue
				//}
			}
		}
		// write remaining part of buffer
		if bn > 0 {
			newN := atomic.AddInt64(&b.Nn, bn) - bn
			for bi = 0; bi < bn; bi++ {
				b.N[newN+int64(bi)] = buffer[bi]
			}
			bn = 0
		}
		doneC <- true
	}
}

// Spawn X processes that all process the current layer in parallel, by
// distributing the work evenly: [0..Cn/X) [Cn/X..2*Cn/X) .. [(X-1)Cn/X..Cn).
// Once they're done, it reports this on the 'done' channel. The main proc
// will wait for everything to finish, swap the current and next queues and
// start again.
func (b *ParBFSLB) Run(g graph.Graph, from int) {
	// init search setup
	b.C[0] = int64(from)
	b.V[from] = 1
	b.Ci = 0 // current queue index
	b.Cn = 1 // current queue length
	b.Ni = 0 // next queue index
	b.Nn = 0 // next queue length
	var stateCount int64 = 0

	doneC := make(chan bool, b.procs)
	rangeC := make([]chan int64, b.procs)

	for p := 0; p < b.procs; p++ {
		rangeC[p] = make(chan int64, 2)
		go b.proc(g, doneC, rangeC[p])
	}

	for b.Cn > 0 {
		step := int(b.Cn) / b.procs
		for p := 0; p < b.procs-1; p++ {
			//fmt.Println("Starting",int64(p*step),"to",int64((p+1)*step+1),"max:",b.Cn)
			rangeC[p] <- int64(p * step)
			rangeC[p] <- int64((p + 1) * step)
		}
		//fmt.Println("Starting",int64((b.procs-1)*step),"to",b.Cn,"max:",b.Cn)
		rangeC[b.procs-1] <- int64((b.procs - 1) * step)
		rangeC[b.procs-1] <- b.Cn

		// wait for all b.procs to finish
		for p := 0; p < b.procs; p++ {
			<-doneC
		}

		//fmt.Println("Finished iteration of", b.Cn, "states")
		stateCount += b.Cn
		b.C, b.N = b.N, b.C
		b.Ci = 0
		b.Ni = 0
		b.Cn = b.Nn
		b.Nn = 0

		time.Sleep(0 * time.Millisecond)
	}

	//fmt.Println("State count and actual", stateCount, g.NumStates(),
	//	"for", b.procs, "procs")
	if stateCount != int64(g.NumStates()) {
		fmt.Println("State count and actual", stateCount, g.NumStates())
		panic("Wrong number of states!")
	}
}

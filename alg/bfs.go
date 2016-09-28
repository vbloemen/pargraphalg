package alg

import (
	//"fmt"

	"github.com/vbloemen/pargraphalg/graph"

	"github.com/karalabe/cookiejar/collections/queue"
)

// Data type for BFS.
type BFS struct {
	Search  // implementing the Search interface
	visited map[int]bool
	q       *queue.Queue
}

// Constructor for the BFS type.
func NewBFS() *BFS {
	q := queue.New()
	return &BFS{visited: make(map[int]bool), q: q}
}

// Performs a BFS.
func (b BFS) Run(g graph.Graph, from int) {
	b.q.Push(from)
	b.visited[from] = true

	for !b.q.Empty() {
		state := b.q.Pop().(int)
		sucs := g.Successors(state)
		for _, si := range sucs {
			if !b.visited[si] {
				b.visited[si] = true
				b.q.Push(si)
			}
		}
	}
}

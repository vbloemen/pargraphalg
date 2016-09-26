package alg

import (
	"fmt"

	"github.com/vbloemen/pargraphalg/graph"
)

// Data type for DFS.
type DFS struct {
	Search  // implementing the Search interface
	visited map[int]bool
}

// Constructor for the DFS type.
func NewDFS() *DFS {
	return &DFS{visited: make(map[int]bool)}
}

// Performs a recursive DFS search.
func (d DFS) Run(g graph.Graph, from int) {
	d.visited[from] = true
	for _, suc := range g.Successors(from) {
		if !d.visited[suc] {
			fmt.Printf("%d -> %d\n", from, suc)
			d.Run(g, suc)
		}
	}
}

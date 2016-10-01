package alg

import (
	//"fmt"

	"github.com/vbloemen/pargraphalg/graph"
)

// Data type for DFS.
type DFS struct {
	Search        // implementing the Search interface
	V      []bool // visited set
}

// Constructor for the DFS type.
func NewDFS() *DFS {
	V := make([]bool, 1e8)
	return &DFS{V: V}
}

// Performs a recursive DFS.
func (d DFS) Run(g graph.Graph, from int) {
	d.V[from] = true
	for _, suc := range g.Successors(from) {
		if !d.V[suc] {
			//fmt.Printf("%d -> %d\n", from, suc)
			d.Run(g, suc)
		}
	}
}

// Package alg provides implementations of graph searching algorithms,
// both sequential and multi-core.
package alg

import (
	"github.com/vbloemen/pargraphalg/graph"
)

// interface object for the Search type.
type Search interface {
	Run(g graph.Graph, from int)
	Init()
}

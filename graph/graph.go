// Package graph provides an interface for specifying graphs and contains
// implementations of several example graphs.
package graph

// Interface object for the graph type.
type Graph interface {
	Init() int
	Successors(state int) []int
}

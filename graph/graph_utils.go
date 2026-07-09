package graph

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Nodes interface {
	constraints.Integer | string | constraints.Float
}

type Graph[T Nodes] struct {
	graph map[string][]string
}

func (g *Graph[T]) InitGraph(graph map[string][]string) *Graph[T] {
	return &Graph[T]{
		graph: graph,
	}
}
func (g *Graph[T]) PrintGraph() {
	fmt.Println(g.graph)
}

func (g *Graph[T]) Length() int {
	return len(g.graph)
}

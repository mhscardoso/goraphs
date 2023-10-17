package alist

import (
	"fmt"

	"github.com/mhscardoso/goraphs/node"
)

type List struct {
	Vector   []*node.Node
	Vertices int
	Edges    int
}

// Allocate memory enough for a list
func (l *List) Allocate(vertices int) {
	vector := make([]*node.Node, vertices)

	l.Vector = vector
	l.Vertices = vertices
}

func (l *List) AddNeighbor(vertex, neighbor int) *node.Node {
	newNeighbor := new(node.Node)
	newNeighbor.Vertex = neighbor
	newNeighbor.Next = l.Vector[vertex]

	l.Vector[vertex] = newNeighbor

	return newNeighbor
}

/* Just checks if a given vertex v already
 * has a neighbor n. Return true case not.
 */
func (l *List) CheckNeighbors(vertex, neighbor int) bool {
	n := l.Neighbors(vertex)
	for ; n != nil; n = n.Next {
		if n.Vertex == neighbor {
			return false
		}
	}

	return true
}

/* Relate two vertices called vertex and neighbor
 * checking if that neighbor was already inserted
 * in the given vertex.
 */
func (l *List) Relate(vertex, neighbor int, edges *int) {
	v := vertex - 1
	n := neighbor - 1

	if !l.CheckNeighbors(v, n) || v == n {
		return
	}

	l.AddNeighbor(v, n)
	l.AddNeighbor(n, v)

	*edges++
}

// Update number of edges of list while
// Read the file
func (l *List) UpdateEdges(edges int) {
	l.Edges = edges
}

// Return all neighbors of list
func (l *List) Neighbors(vertex int) *node.Node {
	return l.Vector[vertex]
}

// Getters of fields Vertices and Edges
func (l List) N() int {
	return l.Vertices
}

func (l List) M() int {
	return l.Edges
}

// Method to test in smalls graphs
func (l *List) See() {
	for i, v := range l.Vector {
		fmt.Printf("V: %v -- ", i+1)

		for e := v; e != nil; e = e.Next {
			fmt.Printf("%v  ,  ", e.Vertex+1)
		}

		fmt.Printf("\n")
	}
}

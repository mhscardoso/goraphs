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

// Relate two vertices in order
func (l *List) Relate(vertex, neighbor int, edges *int) {
	v := vertex - 1
	n := neighbor - 1

	vertexNode := node.NewNode(v)
	neighborNode := node.NewNode(n)

	next_v := l.Vector[v]
	if next_v == nil || next_v.Vertex > n {
		neighborNode.Next = next_v
		l.Vector[v] = neighborNode
	} else {
		for ; next_v.Next != nil; next_v = next_v.Next {
			if next_v.Vertex == n {
				return
			} else if next_v.Next.Vertex > n {
				break
			}
		}

		neighborNode.Next = next_v.Next
		next_v.Next = neighborNode
	}

	next_n := l.Vector[n]
	if next_n == nil || next_n.Vertex > v {
		vertexNode.Next = next_n
		l.Vector[n] = vertexNode
	} else {
		for ; next_n.Next != nil; next_n = next_n.Next {
			if next_n.Vertex == v {
				return
			} else if next_n.Next.Vertex > v {
				break
			}
		}

		vertexNode.Next = next_n.Next
		next_n.Next = vertexNode
	}

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

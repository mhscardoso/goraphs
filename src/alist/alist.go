package alist

import (
	"fmt"

	"github.com/mhscardoso/goraphs/set"
)

type List struct {
	Vector   []set.Set[int]
	Vertices int
	Edges    int
}

// Allocate memory enough for a list
func (l *List) Allocate(vertices int) {
	vector := make([]set.Set[int], vertices)

	for i := range vector {
		vector[i] = make(set.Set[int])
	}

	l.Vector = vector
	l.Vertices = vertices
}

/* Relate two vertices called vertex and neighbor
 * checking if that neighbor was already inserted
 * in the given vertex.
 */
func (l *List) Relate(vertex, neighbor int, edges *int) {
	v := vertex - 1
	n := neighbor - 1

	if v == n || l.Vector[v].Contains(n) {
		return
	}

	l.Vector[v].Add(n)
	l.Vector[n].Add(v)

	*edges++
}

// Update number of edges of list while
// Read the file
func (l *List) UpdateEdges(edges int) {
	l.Edges = edges
}

// Return all neighbors of list
func (l *List) Neighbors(vertex int) set.Set[int] {
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

		for e := range v {
			fmt.Printf("%v  ,  ", e+1)
		}

		fmt.Printf("\n")
	}
}

package awlists

import (
	"fmt"

	"github.com/mhscardoso/goraphs/container/set"
)

type WList[T float64 | int] struct {
	Vector   []set.SetW[int, T]
	Vertices int
	Edges    int
	Targeted bool
}

// Allocate memory enough for a list
func (l *WList[T]) Allocate(vertices int) {
	vector := make([]set.SetW[int, T], vertices)

	for i := range vector {
		vector[i] = make(set.SetW[int, T])
	}

	l.Vector = vector
	l.Vertices = vertices
}

/* Relate two vertices called vertex and neighbor
 * checking if that neighbor was already inserted
 * in the given vertex and weigth.
 */
func (l *WList[T]) Relate(vertex, neighbor int, weigth T, edges *int) {
	v := vertex - 1
	n := neighbor - 1

	if v == n || l.Vector[v].Contains(n) {
		return
	}

	l.Vector[v].Add(n, weigth)

	if !l.Targeted {
		l.Vector[n].Add(v, weigth)
	}

	*edges++
}

// Update number of edges of list while
// Read the file
func (l *WList[T]) UpdateEdges(edges int) {
	l.Edges = edges
}

// Return all neighbors of list
func (l *WList[T]) Neighbors(vertex int) any {
	return l.Vector[vertex]
}

// Getters of fields Vertices and Edges
func (l WList[T]) N() int {
	return l.Vertices
}

func (l WList[T]) M() int {
	return l.Edges
}

// Method to test in smalls graphs
func (l *WList[T]) See() {
	for i, v := range l.Vector {
		fmt.Printf("V: %v -- ", i+1)

		for k, w := range v {
			fmt.Printf("%v: %v,  ", k+1, w)
		}

		fmt.Printf("\n")
	}
}

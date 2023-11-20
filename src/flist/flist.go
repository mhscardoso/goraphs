package flist

import (
	"fmt"

	"github.com/mhscardoso/goraphs/container/set"
)

// Graphs with weigth and flow
type FList struct {
	Vector   []set.SetF
	Vertices int
	Edges    int
	Targeted bool
}

// Allocate memory enough for a list
func (l *FList) Allocate(vertices int) {
	vector := make([]set.SetF, vertices)

	for i := range vector {
		vector[i] = make(set.SetF)
	}

	l.Vector = vector
	l.Vertices = vertices
}

func (l *FList) Relate(vertex, neighbor int, weigth float64, edges *int) {
	v := vertex - 1
	n := neighbor - 1

	if v == n {
		return
	}

	l.Vector[v].Add(uint32(n), int(weigth), 0)

	if !l.Targeted {
		l.Vector[n].Add(uint32(v), int(weigth), 0)
	}

	*edges++
}

// Update number of edges of list while
// Read the file
func (l *FList) UpdateEdges(edges int) {
	l.Edges = edges
}

// Return all neighbors of list
func (l *FList) Neighbors(vertex int) any {
	return l.Vector[vertex]
}

// Getters of fields Vertices and Edges
func (l FList) N() int {
	return l.Vertices
}

func (l FList) M() int {
	return l.Edges
}

// Method to test in smalls graphs
func (l *FList) See() {
	for i, v := range l.Vector {
		fmt.Printf("V: %v -- ", i+1)

		for k, w := range v {
			fmt.Printf("[N: %v: | W: %v | F: %v],  ", k+1, w[0], w[1])
		}

		fmt.Printf("\n")
	}
}

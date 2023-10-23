package amatrix

import (
	"fmt"

	"github.com/mhscardoso/goraphs/container/set"
)

type Matrix struct {
	G        [][]byte
	Vertices int
	Edges    int
}

func (matrix *Matrix) Allocate(vertices int) {
	matrix.Vertices = vertices
	matrix.G = make([][]byte, vertices)
	lines := make([]byte, vertices*vertices)

	for i := range matrix.G {
		matrix.G[i], lines = lines[:vertices], lines[vertices:]
	}
}

func (m *Matrix) Relate(vertex, neighbor int, edges *int) {
	if vertex == neighbor {
		return
	}

	v := vertex - 1
	n := neighbor - 1

	if m.G[v][n] == 1 {
		return
	}

	m.G[v][n] = 1
	m.G[n][v] = 1

	*edges++
}

func (matrix *Matrix) Neighbors(vertex int) set.Set[int] {
	s := make(set.Set[int])

	for i := range matrix.G[vertex] {
		if matrix.G[vertex][i] == 1 {
			s.Add(i)
		}
	}

	return s
}

func (matrix *Matrix) UpdateEdges(edges int) {
	matrix.Edges = edges
}

func (matrix Matrix) N() int {
	return matrix.Vertices
}

func (matrix Matrix) M() int {
	return matrix.Edges
}

func (m *Matrix) See() {
	for i := 0; i < m.Vertices; i++ {
		for j := 0; j < m.Vertices; j++ {
			fmt.Printf("%v ", m.G[i][j])
		}
		fmt.Printf("\n")
	}
}

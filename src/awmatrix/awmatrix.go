package awmatrix

import (
	"fmt"

	"github.com/mhscardoso/goraphs/container/set"
)

type WMatrix struct {
	G        [][]float32
	Vertices int
	Edges    int
	Targeted bool
}

func (matrix *WMatrix) Allocate(vertices int) {
	matrix.Vertices = vertices
	matrix.G = make([][]float32, vertices)
	lines := make([]float32, vertices*vertices)

	for i := range matrix.G {
		matrix.G[i], lines = lines[:vertices], lines[vertices:]
	}
}

func (m *WMatrix) Relate(vertex, neighbor int, weigth float64, edges *int) {
	if vertex == neighbor {
		return
	}

	v := vertex - 1
	n := neighbor - 1

	if m.G[v][n] != 0 {
		return
	}

	m.G[v][n] = float32(weigth)

	if !m.Targeted {
		m.G[n][v] = float32(weigth)
	}

	*edges++
}

func (matrix *WMatrix) Neighbors(vertex int) any {
	s := make(set.SetW[int, float64])

	for i := range matrix.G[vertex] {
		if matrix.G[vertex][i] != 0 {
			s.Add(i, float64(matrix.G[vertex][i]))
		}
	}

	return s
}

func (matrix *WMatrix) UpdateEdges(edges int) {
	matrix.Edges = edges
}

func (matrix WMatrix) N() int {
	return matrix.Vertices
}

func (matrix WMatrix) M() int {
	return matrix.Edges
}

func (m *WMatrix) See() {
	for i := 0; i < m.Vertices; i++ {
		for j := 0; j < m.Vertices; j++ {
			fmt.Printf("%v ", m.G[i][j])
		}
		fmt.Printf("\n")
	}
}

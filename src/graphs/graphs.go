package graphs

import (
	"github.com/mhscardoso/goraphs/src/alist"
	"github.com/mhscardoso/goraphs/src/amatrix"
	"github.com/mhscardoso/goraphs/src/awlists"
	"github.com/mhscardoso/goraphs/src/awmatrix"
	"github.com/mhscardoso/goraphs/src/flist"
)

type Graph interface {
	Neighbors(vertex int) any
	Relate(vertex, neighbor int, weigth float64, edges *int)
	Allocate(vertices int)
	UpdateEdges(edges int)
	See()
	N() int
	M() int
}

// Constructors
func List(targeted bool) *alist.List {
	initial := new(alist.List)
	initial.Targeted = targeted

	return initial
}

func Matrix(targeted bool) *amatrix.Matrix {
	initial := new(amatrix.Matrix)
	initial.Targeted = targeted

	return initial
}

// Constructors for graphs with weigths
func WList(targeted bool) *awlists.WList {
	initial := new(awlists.WList)
	initial.Targeted = targeted

	return initial
}

func WMatrix(targeted bool) *awmatrix.WMatrix {
	initial := new(awmatrix.WMatrix)
	initial.Targeted = targeted

	return initial
}

// Constructor for FList
func FList(targeted bool) *flist.FList {
	initial := new(flist.FList)
	initial.Targeted = targeted

	return initial
}

func MyWay(parents []int, vertex int) []int {
	if vertex < 1 || vertex >= len(parents) {
		return nil
	}

	way := make([]int, 0, len(parents))

	way = append(way, vertex)

	parent := parents[vertex-1]
	for parent != 0 {
		way = append(way, parent)
		parent = parents[parent-1]
	}

	return way
}

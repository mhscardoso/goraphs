package graphs

import (
	"github.com/mhscardoso/goraphs/src/alist"
	"github.com/mhscardoso/goraphs/src/amatrix"
	"github.com/mhscardoso/goraphs/src/awlists"
	"github.com/mhscardoso/goraphs/src/awmatrix"
)

type Graph interface {
	Neighbors(vertex int) any
	Relate(vertex, neighbor int, weigth float32, edges *int)
	Allocate(vertices int)
	UpdateEdges(edges int)
	See()
	N() int
	M() int
}

// Constructors
func List() *alist.List {
	return new(alist.List)
}

func Matrix() *amatrix.Matrix {
	return new(amatrix.Matrix)
}

// Constructors for graphs with weigths
func WList() *awlists.WList {
	return new(awlists.WList)
}

func WMatrix() *awmatrix.WMatrix {
	return new(awmatrix.WMatrix)
}

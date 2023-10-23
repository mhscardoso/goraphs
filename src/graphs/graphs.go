package graphs

import (
	"github.com/mhscardoso/goraphs/container/set"
	"github.com/mhscardoso/goraphs/src/alist"
	"github.com/mhscardoso/goraphs/src/amatrix"
)

type Graph interface {
	Neighbors(vertex int) set.Set[int]
	Relate(vertex, neighbor int, edges *int)
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

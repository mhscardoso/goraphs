package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/list"
	"github.com/mhscardoso/goraphs/matrix"
)

func main() {

	fmt.Printf("Lendo o Grafo do disco...\n")

	A, err := matrix.CreateMatrix("../grafos/grafo_2.txt")
	if err != nil {
		panic(err)
	}

	B, err := list.CreateList("../grafos/grafo_2.txt", list.AddSorted)
	if err != nil {
		panic(err)
	}

	// N, M, min, max, median, middle := A.GetInfo()
	// fmt.Printf("Vertices: %v\n", N)
	// fmt.Printf("Edges: %v\n", M)
	// fmt.Printf("Min: %v\n", min)
	// fmt.Printf("Max: %v\n", max)
	// fmt.Printf("Median: %v\n", median)
	// fmt.Printf("Middle: %v\n", middle)

	_, _, _, _, _, _, edgesA := A.GetInfo()
	_, _, _, _, _, _, edgesB := B.GetInfo()

	for i := range edgesA {
		if edgesA[i] != edgesB[i] {
			fmt.Printf("i: %v, A: %v, B: %v\n", i, edgesA[i], edgesB[i])
		}
	}

}

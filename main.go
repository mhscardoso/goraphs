package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/list"
	//"github.com/mhscardoso/goraphs/matrix"
)

func main() {

	fmt.Printf("Lendo o Grafo do disco...\n")

	// A, err := matrix.CreateMatrix("../grafos/grafo_1.txt")
	// if err != nil {
	// 	panic(err)
	// }

	B, err := list.CreateList("../grafos/grafo_3.txt", list.AddInOrder)
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

	B.SaveData("list_data/grafo3aaa_List.txt")
	// _, parent, level := B.BFS(10)

	// for i := range parent {
	// 	if parent[i] == 0 {
	// 		fmt.Printf("V: %v, P: %v, L: %v\n", i+1, parent[i], level[i])
	// 	}
	// }

}

package main

import (
	"fmt"

	//"github.com/mhscardoso/goraphs/matrix"
	"github.com/mhscardoso/goraphs/list"
)

func main() {

	fmt.Printf("Lendo o Grafo do disco...\n")

	A, err := list.CreateList("../grafos/grafo_1.txt", list.AddInOrder)
	if err != nil {
		panic(err)
	}

	dist := A.Diameter()
	fmt.Printf("Distancia: %v\n", dist)

}

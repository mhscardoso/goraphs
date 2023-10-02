package main

import (
	"fmt"

	//"github.com/mhscardoso/goraphs/matrix"
	"github.com/mhscardoso/goraphs/list"
)

func main() {

	fmt.Printf("Lendo o Grafo do disco...\n")

	A, err := list.CreateList("../grafos/grafo_6.txt", list.AddInOrder)
	if err != nil {
		panic(err)
	}

	measure := A.StatsBFS()
	fmt.Printf("Tempo medio: %v\n", measure)

}

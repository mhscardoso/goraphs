package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/graphs"
)

func main() {
	A := graphs.List()

	fmt.Printf("Lendo Grafo do Disco...\n")
	graphs.ReadFile(A, "../grafos/grafo_2.txt")

	fmt.Printf("Grafo Lido\n")

	components, _ := graphs.ConectedComponents(A)

	for _, comp := range components {
		fmt.Printf("%v\n", comp.Length)
	}
}

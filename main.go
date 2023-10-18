package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/src/graphs"
)

func main() {
	A := graphs.Matrix()

	fmt.Printf("Lendo Grafo do Disco...\n")
	graphs.ReadFile(A, "../grafos/grafo_2.txt")

	fmt.Printf("Grafo Lido\n")

	c := graphs.ConectedComponents(A)

	c.See()
}

package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/src/graphs"
)

func main() {
	A := graphs.List()

	fmt.Printf("Lendo Grafo do Disco...\n")
	graphs.ReadFile(A, "../grafos/grafo_1.txt")

	fmt.Printf("Grafo Lido\n")

	d := graphs.Distance(A, 10, 400)
	fmt.Printf("D: %v\n", d)
}

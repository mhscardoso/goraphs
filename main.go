package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/src/graphs"
)

func main() {
	A := graphs.WMatrix()

	fmt.Printf("Lendo Grafo do Disco...\n")
	graphs.ReadFile(A, "../grafos/TP2/grafo_W_1.txt")

	fmt.Printf("Grafo Lido\n")

	// graphs.GetInfo(A)

	// A.See()

	graphs.Dijkstra(A, 3)
}

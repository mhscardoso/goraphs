package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/src/graphs"
)

func main() {
	A := graphs.WList()

	fmt.Printf("Lendo Grafo do Disco...\n")
	graphs.ReadFile(A, "../grafos/TP2/grafo_W_2.txt")

	fmt.Printf("Grafo Lido\n")

	res := graphs.Dijkstra(A, 10)

	graphs.WriteResults("grafo_2.txt", res, 10)
}

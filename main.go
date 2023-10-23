package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/src/graphs"
)

func main() {
	A := graphs.List()

	fmt.Printf("Lendo Grafo do Disco...\n")
	graphs.ReadFile(A, "../grafos/TP1/grafo_6.txt")

	fmt.Printf("Grafo Lido\n")

	graphs.GetInfo(A)
}

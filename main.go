package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/graphs"
)

func main() {
	A := graphs.List()

	fmt.Printf("Lendo Grafo do Disco...\n")
	graphs.ReadFile(A, "../grafos/grafo_3.txt")

	fmt.Printf("Grafo Lido\n")

	graphs.GetInfo(A)
}

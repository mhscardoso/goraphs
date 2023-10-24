package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/src/graphs"
)

func main() {
	A := graphs.WList()

	fmt.Printf("Lendo Grafo do Disco...\n")
	graphs.ReadFile(A, "testw.txt")

	fmt.Printf("Grafo Lido\n")

	graphs.GetInfo(A)

	A.See()
}

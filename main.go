package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/container/heap"
)

func main() {
	// A := graphs.WList()

	// fmt.Printf("Lendo Grafo do Disco...\n")
	// graphs.ReadFile(A, "../grafos/TP2/grafo_W_5.txt")

	// fmt.Printf("Grafo Lido\n")

	// graphs.GetInfo(A)

	// A.See()

	h := heap.New(10)
	h.Insert(1, 1)
	h.Insert(2, 3)
	h.Insert(3, 6)
	h.Insert(4, 5)
	h.Insert(5, 9)
	h.Insert(6, 8)

	fmt.Printf("%v\n\n", h)

	h.Remove()

	fmt.Printf("%v\n", h)

	h.Remove()

	fmt.Printf("%v\n", h)

	h.Remove()

	fmt.Printf("%v\n", h)

	h.Remove()

	fmt.Printf("%v\n", h)
}

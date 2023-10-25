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
	h.Insert(10, 1)
	h.Insert(2, 2)
	h.Insert(7, 3)
	h.Insert(5, 4)
	h.Insert(4, 5)
	h.Insert(8, 6)

	fmt.Printf("%v\n", h)

	h.Remove()

	fmt.Printf("%v\n", h)

	h.Remove()

	fmt.Printf("%v\n", h)

	h.Remove()

	fmt.Printf("%v\n", h)

	h.Remove()

	fmt.Printf("%v\n", h)
}

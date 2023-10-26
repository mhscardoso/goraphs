package main

import (
	"fmt"

	//"github.com/mhscardoso/goraphs/container/heap"
	"github.com/mhscardoso/goraphs/src/graphs"
)

func main() {
	// hp := heap.New(10)

	// hp.Update(4, 10)
	// hp.Update(1, 1)
	// hp.Update(2, 3)
	// hp.Update(6, 4)
	// hp.Update(7, 7)
	// hp.Update(8, 2)
	// hp.Update(0, 9)
	// hp.Update(3, 2)
	// hp.Update(5, 6)
	// hp.Update(9, 3)

	// fmt.Printf("%v\n", hp)

	// hp.Remove()
	// hp.Remove()
	// hp.Remove()

	// fmt.Printf("%v\n", hp)

	A := graphs.WList()

	graphs.ReadFile(A, "../grafos/TP2/grafo_W_1.txt")

	fmt.Printf("Grafo Lido\n")

	_, parent := graphs.Dijkstra(A, 10)

	fmt.Printf("DKT Feito\n")

	//graphs.WriteResults("results/heaps/grafo_W_5.txt", dist, parent, 10)

	way := graphs.MyWay(parent, 20)

	fmt.Printf("%v\n", way)
}

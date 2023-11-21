package main

import (
	"fmt"
	"time"

	"github.com/mhscardoso/goraphs/src/graphs"
)

func main() {
	A := graphs.FList(true)

	fmt.Printf("Reading...\n")
	graphs.ReadFileIntFlow(A, "../grafos/TP3/grafo_rf_1.txt")
	//graphs.ReadFileIntFlow(A, "testf.txt")

	fmt.Printf("Calculating...\n")

	t1 := time.Now()

	graphs.FordFulkerson(A, 1, 2, 512)

	t2 := time.Now()

	fmt.Printf("Done.\n")

	ti := t2.Sub(t1).Seconds()

	f1 := 0

	for _, v := range A.Vector[0] {
		f1 += v[1]
	}

	fmt.Printf("Exit flow: %v\n", f1)

	fmt.Printf("Execution Time: %v\n", ti)
}

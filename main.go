package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/graphs"
)

func main() {
	A := graphs.Matrix()

	graphs.ReadFile(A, "test.txt")

	A.See()

	_, parent, _ := graphs.BFS(A, 1)
	for i, v := range parent {
		fmt.Printf("%v - %v\n", i+1, v)
	}
}

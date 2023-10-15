package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/graphs"
)

func main() {
	A := graphs.List()

	graphs.ReadFile(A, "../grafos/grafo_6.txt")

	components, _ := graphs.ConectedComponents(A)

	for _, comp := range components {
		fmt.Printf("%v\n", comp.Length)
	}
}

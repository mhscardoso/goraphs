package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/list"
	//"github.com/mhscardoso/goraphs/matrix"
)

func main() {
	fmt.Printf("Lendo o Grafo do disco...\n")

	fu := list.AddInOrder

	ls, err := list.CreateList("test.txt", fu)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Mostrando!\n")

	ls.See()

	fmt.Printf("Grafo lido, fazendo a BFS\n")

	parent, level := ls.DFS(1)

	for i, v := range parent {
		fmt.Printf("V: %v -- P: %v -- L: %v\n", i+1, v, level[i])
	}
}

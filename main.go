package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/list"
	//"github.com/mhscardoso/goraphs/matrix"
)

func main() {
	fmt.Printf("Lendo o Grafo do disco...\n")
	ls, err := list.CreateList("../grafos/grafo_6.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Grafo lido, fazendo a BFS\n")

	parent, level := ls.BFS(10)

	for i, v := range parent {
		fmt.Printf("V: %v -- P: %v -- L: %v\n", i+1, v, level[i])
	}
}

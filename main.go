package main

import (
	"github.com/mhscardoso/goraphs/list"
)

func main() {
	ls, err := list.CreateMatrix("../grafos/grafo_1.txt")
	if err != nil {
		panic(err)
	}

	ls.See()
}

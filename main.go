package main

import (
	"github.com/mhscardoso/goraphs/src/graphs"
)

func main() {
	A := graphs.FList(true)

	graph := "grafo_rf_3.txt"

	graph_path := "../grafos/TP3/" + graph
	graphs.ReadFile(A, graph_path)

	results_path := "results/" + graph
	graphs.WriteFordResult(A, 1, 2, 1024, results_path)
}

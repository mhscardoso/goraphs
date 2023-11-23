package graphs

import (
	"fmt"
	"testing"
	"time"
)

// Corretor, edite estas variáveis para que possa ler os arquivos corretos!
var filename = []string{
	"../../../grafos/TP3/grafo_rf_1.txt",
	"../../../grafos/TP3/grafo_rf_2.txt",
	"../../../grafos/TP3/grafo_rf_3.txt",
	"../../../grafos/TP3/grafo_rf_4.txt",
	"../../../grafos/TP3/grafo_rf_5.txt",
	"../../../grafos/TP3/grafo_rf_6.txt",
}

func TestFord(t *testing.T) {
	rounds := 10
	source := 1
	target := 2
	delta := 1

	for _, name := range filename {
		var totalTime float64 = 0

		fmt.Printf("----------------------- Arquivo %v\n", name)
		for i := 0; i < rounds; i++ {
			A := FList(true)
			ReadFile(A, name)

			fmt.Printf("%vº Ford-Fulkerson - ", i+1)

			t1 := time.Now()
			FordFulkerson(A, source, target, delta)
			t2 := time.Now()

			totalTime += t2.Sub(t1).Seconds()

			// Calculando Fluxo Máximo
			maxFlow := 0
			for _, v := range A.Vector[source-1] {
				maxFlow += v.GetFlow()
			}

			fmt.Printf("Fluxo Máximo entre %v e %v: %v\n", source, target, maxFlow)
		}

		fmt.Printf("\nArquivo Lido: %v\n", name)
		fmt.Printf("Tempo médio: %v\n", totalTime/float64(rounds))
		fmt.Printf("------------------------------\n\n")
	}
}

func TestConnectedFlow(t *testing.T) {
	g_path := filename[5]

	A := List(true)

	fmt.Printf("Lendo Arquivo: %v\n", g_path)
	ReadFile(A, g_path)

	fmt.Printf("Obtendo Componentes Conexas...\n\n")
	q := ConectedComponents(A)

	fmt.Printf("Componentes conexas para %v\n", g_path)
	for e := q.First; e != nil; e = e.Next {
		fmt.Printf("%v\n", e.Vertex.Length)
	}
}

func TestBFSFord(t *testing.T) {
	g_path := filename[4]
	delta := 16384
	source := 1
	target := 2

	A := FList(true)

	fmt.Printf("Lendo Arquivo: %v\n", g_path)
	ReadFile(A, g_path)

	fmt.Printf("Ford-Fulkerson...\n")
	times, _, _, _, _ := FordFulkersonStats(A, source, target, delta)

	fmt.Printf("Informações para: %v\n", g_path)
	fmt.Printf("Delta: %v\n", delta)
	fmt.Printf("Número de BFS: %v\n", times)
}

func TestGraphDensity(t *testing.T) {
	g_path := filename[5]

	A := FList(true)

	fmt.Printf("Lendo Arquivo: %v\n", g_path)
	ReadFile(A, g_path)

	edges := A.Edges
	vertices := A.Vertices

	maxEdges := ((vertices - 1) * vertices) / 2
	density := float64(edges) / float64(maxEdges)

	fmt.Printf("Graph: %v\n", g_path)
	fmt.Printf("Density: %.12f\n", density)
}

func TestStatsFord(t *testing.T) {
	g_path := filename[5]
	delta := 1
	source := 1
	target := 2

	A := FList(true)

	fmt.Printf("Lendo Arquivo: %v\n", g_path)
	ReadFile(A, g_path)

	fmt.Printf("Ford-Fulkerson...\n")
	BFSTimes, BFSTotal, residualTime, augs, augTime := FordFulkersonStats(A, source, target, delta)

	BFSMedia := BFSTotal / float64(BFSTimes)
	AugsMedia := augTime / float64(augs)

	fmt.Printf("Informações para: %v\n", g_path)
	fmt.Printf("Delta: %v\n", delta)
	fmt.Printf("Número de BFS: %v\n", BFSTimes)
	fmt.Printf("Tempo para o grafo residual (1º vez): %v\n", residualTime)
	fmt.Printf("Média para BFSs: %v\n", BFSMedia)
	fmt.Printf("Augments: %v\n", augs)
	fmt.Printf("Média para AugmentResidual: %v\n", AugsMedia)
}

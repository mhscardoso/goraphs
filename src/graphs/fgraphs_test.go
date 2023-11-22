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

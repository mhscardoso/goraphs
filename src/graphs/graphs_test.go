package graphs

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Corretor, edite estas variáveis para que possa ler os arquivos corretos!
var filenames = []string{
	"../../../grafos/TP1/grafo_1.txt",
	"../../../grafos/TP1/grafo_2.txt",
	"../../../grafos/TP1/grafo_3.txt",
	"../../../grafos/TP1/grafo_4.txt",
	"../../../grafos/TP1/grafo_5.txt",
	"../../../grafos/TP1/grafo_6.txt",
}

func TestListCreate(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Lista de Adjacências para %v\n", name)
		A := List()

		t1 := time.Now()
		ReadFile(A, name)

		t2 := time.Now()

		diff := t2.Sub(t1)

		fmt.Printf("Leitura do arquivo e criação da Lista: %v s\n", float64(diff.Seconds()))

		fmt.Printf("Vértices: %v\nArestas: %v\n\n", A.Vertices, A.Edges)
	}
}

func TestListStatsBFS(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Lista de Adjacências para %v\n", name)
		A := List()
		ReadFile(A, name)

		fmt.Printf("Grafo lido como lista\n")

		var times float64 = 0
		s := rand.Intn(A.Vertices)

		for i := 0; i < 100; i++ {
			t1 := time.Now()
			BFS(A, s, 0, nil)
			t2 := time.Now()

			times += t2.Sub(t1).Seconds()
		}

		fmt.Printf("Tempo médio da BFS para %v: %v\n\n", name, times/100)
	}
}

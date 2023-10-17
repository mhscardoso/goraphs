package graphs_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mhscardoso/goraphs/graphs"
)

// Corretor, edite estas variáveis para que possa ler os arquivos corretos!
var filenames = []string{
	"../../grafos/grafo_1.txt",
	"../../grafos/grafo_2.txt",
	"../../grafos/grafo_3.txt",
	"../../grafos/grafo_4.txt",
	"../../grafos/grafo_5.txt",
	"../../grafos/grafo_6.txt",
}

func TestListCreate(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Lista de Adjacências para %v\n", name)
		A := graphs.List()

		t1 := time.Now()
		graphs.ReadFile(A, name)

		t2 := time.Now()

		diff := t2.Sub(t1)

		fmt.Printf("Leitura do arquivo e criação da Lista: %v s\n", float64(diff.Seconds()))

		fmt.Printf("Vértices: %v\nArestas: %v\n\n", A.Vertices, A.Edges)
	}
}

func TestListEdges(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Lista de Adjacências para %v\n", name)
		A := graphs.List()

		graphs.GetInfo(A)
	}
}

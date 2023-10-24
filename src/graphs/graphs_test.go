package graphs

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// Corretor, edite estas variáveis para que possa ler os arquivos corretos!
var filenames = []string{
	"../../../grafos/TP1/grafo_1.txt",
	"../../../grafos/TP1/grafo_2.txt",
	// "../../../grafos/TP1/grafo_3.txt",
	// "../../../grafos/TP1/grafo_4.txt",
	// "../../../grafos/TP1/grafo_5.txt",
	// "../../../grafos/TP1/grafo_6.txt",
}

func TestCreate(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Estrutura para %v\n", name)

		// Edit this line to read
		// like a matrix or a list.
		// A := Matrix() -- like Matrix
		// A := List()   -- like List
		A := List()

		fmt.Printf("Grafo lido como %v\n", reflect.TypeOf(*A))

		t1 := time.Now()
		ReadFile(A, name)

		t2 := time.Now()

		diff := t2.Sub(t1)

		fmt.Printf("Leitura do arquivo e criação de %v: %v s\n",
			reflect.TypeOf(*A), float64(diff.Seconds()))

		fmt.Printf("Vértices: %v\nArestas: %v\n\n", A.Vertices, A.Edges)
	}
}

func TestStatsBFS(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Estrutura para %v\n", name)

		// Edit this line to read
		// like a matrix or a list.
		// A := Matrix() -- like Matrix
		// A := List()   -- like List
		A := List()

		ReadFile(A, name)

		fmt.Printf("Grafo lido como %v\n", reflect.TypeOf(*A))

		var times float64 = 0

		for i := 0; i < 100; i++ {
			s := rand.Intn(A.Vertices)

			t1 := time.Now()
			BFS(A, s, 0, nil)
			t2 := time.Now()

			times += t2.Sub(t1).Seconds()
		}

		fmt.Printf("Tempo médio da BFS para %v: %v\n\n", name, times/100)
	}
}

func TestStatsDFS(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Estrutura para %v\n", name)

		// Edit this line to read
		// like a matrix or a list.
		// A := Matrix() -- like Matrix
		// A := List()   -- like List
		A := List()

		ReadFile(A, name)

		fmt.Printf("Grafo lido como %v\n", reflect.TypeOf(*A))

		var times float64 = 0

		for i := 0; i < 100; i++ {
			s := rand.Intn(A.Vertices)

			t1 := time.Now()
			DFS(A, s, nil)
			t2 := time.Now()

			times += t2.Sub(t1).Seconds()
		}

		fmt.Printf("Tempo médio da BFS para %v: %v\n\n", name, times/100)
	}
}

func TestParents(t *testing.T) {
	testing_vertices := [3]int{1, 2, 3}

	for _, name := range filenames {
		fmt.Printf("Criando Estrutura para %v\n", name)

		// Edit this line to read
		// like a matrix or a list.
		// A := Matrix() -- like Matrix
		// A := List()   -- like List
		A := Matrix()

		ReadFile(A, name)

		fmt.Printf("Grafo lido como %v\n", reflect.TypeOf(*A))

		for _, v := range testing_vertices {
			parent, _ := BFS(A, v, 0, nil)
			fmt.Printf("Raiz: %v - Vértice: %v - Pai: %v\n", v, 10, parent[9])
			fmt.Printf("Raiz: %v - Vértice: %v - Pai: %v\n", v, 20, parent[19])
			fmt.Printf("Raiz: %v - Vértice: %v - Pai: %v\n", v, 30, parent[29])
		}

		fmt.Println()
	}
}

func TestDistances(t *testing.T) {
	testing_vertices := [3]int{10, 20, 30}

	for _, name := range filenames {
		fmt.Printf("Criando Estrutura para %v\n", name)

		// Edit this line to read
		// like a matrix or a list.
		// A := Matrix() -- like Matrix
		// A := List()   -- like List
		A := List()

		ReadFile(A, name)

		fmt.Printf("Grafo lido como %v\n", reflect.TypeOf(*A))

		for i := 0; i < 3; i++ {
			for j := i + 1; j < 3; j++ {
				dist := Distance(A, testing_vertices[i], testing_vertices[j])
				fmt.Printf("(%v, %v) - %v\n", testing_vertices[i], testing_vertices[j], dist)
			}
		}
		fmt.Println()
	}
}

func TestConnectedComponents(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Estrutura para %v\n", name)

		// Edit this line to read
		// like a matrix or a list.
		// A := Matrix() -- like Matrix
		// A := List()   -- like List
		A := List()

		ReadFile(A, name)

		fmt.Printf("Grafo lido como %v\n", reflect.TypeOf(*A))

		components := ConectedComponents(A)

		fmt.Printf("Grafo de %v\nComponentes Conexas: %v\n",
			name, components.Length)

		fmt.Printf("Components:\n")
		for e := components.First; e != nil; e = e.Next {
			fmt.Printf("%v\n", e.Vertex)
		}

		fmt.Printf("\n\n")
	}
}

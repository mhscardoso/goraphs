package list

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Corretor, edite estas variáveis para que possa ler os arquivos corretos!
var filenames = [6]string{
	"../../grafos/grafo_1.txt",
	"../../grafos/grafo_2.txt",
	"../../grafos/grafo_3.txt",
	"../../grafos/grafo_4.txt",
	"../../grafos/grafo_5.txt",
	"../../grafos/grafo_6.txt",
}

func TestCreateList(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Lista de Adjacências para %v\n", name)

		t1 := time.Now()
		L, err := CreateList(name, AddInOrder)
		if err != nil {
			t.Fatalf("Error %v in reading file.", err)
		}
		t2 := time.Now()

		diff := t2.Sub(t1)

		fmt.Printf("Leitura do arquivo e criação da Lista: %v\n", float64(diff.Seconds()))

		fmt.Printf("Vértices: %v\nArestas: %v\n", L.N, L.M)
	}
}

func TestMeasureBFS(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Lista de Adjacências\n")

		L, err := CreateList(name, AddInOrder)
		if err != nil {
			t.Fatalf("Error %v in reading file.", err)
		}

		fmt.Printf("Lista Criada\n")

		var time float64 = 0
		for i := 1; i <= 100; i++ {
			s := rand.Intn(L.N)
			time += float64(L.MeasureBFS(s))
		}

		fmt.Printf("Testando 100 BFSs para %v\n", name)
		fmt.Printf("Arquivo de entrada: %v\n", name)
		fmt.Printf("100 BFSs: %v\n", time)
		fmt.Printf("Tempo médio: %v\n\n", time/float64(100))
	}
}

func TestMeasureDFS(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Lista de Adjacências\n")

		L, err := CreateList(name, AddInOrder)
		if err != nil {
			t.Fatalf("Error %v in reading file.", err)
		}

		fmt.Printf("Lista Criada\n")

		var time float64 = 0
		for i := 1; i <= 100; i++ {
			s := rand.Intn(L.N)
			time += float64(L.MeasureDFS(s))
		}

		fmt.Printf("Testando 100 DFSs para %v\n", name)
		fmt.Printf("Arquivo de entrada: %v\n", name)
		fmt.Printf("100 DFSs: %v\n", time)
		fmt.Printf("Tempo médio: %v\n\n", time/float64(100))
	}
}

func TestParents(t *testing.T) {
	testing_vertices := [3]int{1, 2, 3}

	for _, name := range filenames {
		fmt.Printf("Criando Lista de Adjacências\n")

		L, err := CreateList(name, AddInOrder)
		if err != nil {
			t.Fatalf("Error %v in reading file.", err)
		}

		fmt.Printf("Lista Criada\n")

		fmt.Printf("Arquivo de entrada: %v\n", name)

		for _, v := range testing_vertices {
			_, parent, _ := L.BFS(v)
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
		fmt.Printf("Criando Lista de Adjacências\n")

		L, err := CreateList(name, AddInOrder)
		if err != nil {
			t.Fatalf("Error %v in reading file.", err)
		}

		fmt.Printf("Lista Criada\n")

		fmt.Printf("Arquivo de entrada: %v\n", name)
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 3; j++ {
				dist := L.Distance(testing_vertices[i], testing_vertices[j])
				fmt.Printf("(%v, %v) - %v\n", testing_vertices[i], testing_vertices[j], dist)
			}
		}
		fmt.Println()
	}
}

func TestConnectedComponents(t *testing.T) {
	for _, name := range filenames {
		fmt.Printf("Criando Lista de Adjacências\n")

		L, err := CreateList(name, AddInOrder)
		if err != nil {
			t.Fatalf("Error %v in reading file.", err)
		}

		fmt.Printf("Lista Criada\n")

		components, graphs := L.ConectedComponents()

		fmt.Printf("Grafo de %v\nComponentes Conexas: %v\nMaior Componente: %v\nMenor Componente: %v\n\n",
			name, graphs, components[graphs-1].Size, components[0].Size)
	}
}

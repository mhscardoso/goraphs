package amatrix

// import (
// 	"fmt"
// 	"math/rand"
// 	"testing"
// 	"time"
// )

// // Corretor, edite estas variáveis para que possa ler os arquivos corretos!

// // Em minha máquina, os grafos 5 e 6 não foram suportados utilizando Matrizes
// var filenames = []string{
// 	"../../grafos/grafo_1.txt",
// 	"../../grafos/grafo_2.txt",
// 	"../../grafos/grafo_3.txt",
// 	"../../grafos/grafo_4.txt",
// }

// func TestCreateMatrix(t *testing.T) {
// 	for _, name := range filenames {
// 		fmt.Printf("Criando Matriz de Adjacências para %v\n", name)

// 		t1 := time.Now()
// 		A, err := CreateMatrix(name)
// 		if err != nil {
// 			t.Fatalf("Error %v in reading file.", err)
// 		}
// 		t2 := time.Now()

// 		diff := t2.Sub(t1)

// 		fmt.Printf("Leitura do arquivo e criação da Matriz: %v s\n", float64(diff.Seconds()))

// 		time.Sleep(10 * time.Second)

// 		fmt.Printf("Vértices: %v\nArestas: %v\n\n", A.N, A.M)
// 	}
// }

// func TestMeasureBFS(t *testing.T) {
// 	for _, name := range filenames {
// 		fmt.Printf("Criando Matriz de Adjacências\n")

// 		A, err := CreateMatrix(name)
// 		if err != nil {
// 			t.Fatalf("Error %v in reading file.", err)
// 		}

// 		fmt.Printf("Matriz Criada\n")

// 		var time float64 = 0
// 		for i := 1; i <= 100; i++ {
// 			s := rand.Intn(A.N)
// 			time += float64(A.MeasureBFS(s))
// 			fmt.Printf("Matriz: %v - Tempo: %v\n", i, time)
// 		}

// 		fmt.Printf("Testando 100 BFSs para %v\n", name)
// 		fmt.Printf("Arquivo de entrada: %v\n", name)
// 		fmt.Printf("100 BFSs: %v\n", time)
// 		fmt.Printf("Tempo médio: %v\n\n", time/float64(100))
// 	}
// }

// func TestMeasureDFS(t *testing.T) {
// 	for _, name := range filenames {
// 		fmt.Printf("Criando Matriz de Adjacências\n")

// 		A, err := CreateMatrix(name)
// 		if err != nil {
// 			t.Fatalf("Error %v in reading file.", err)
// 		}

// 		fmt.Printf("Matriz Criada\n")

// 		var time float64 = 0
// 		for i := 1; i <= 100; i++ {
// 			s := rand.Intn(A.N)
// 			time += float64(A.MeasureDFS(s))
// 			fmt.Printf("Matriz: %v - Tempo: %v\n", i, time)
// 		}

// 		fmt.Printf("Testando 100 DFSs para %v\n", name)
// 		fmt.Printf("Arquivo de entrada: %v\n", name)
// 		fmt.Printf("100 DFSs: %v\n", time)
// 		fmt.Printf("Tempo médio: %v\n\n", time/float64(100))
// 	}
// }

// func TestParents(t *testing.T) {
// 	testing_vertices := [3]int{1, 2, 3}

// 	for _, name := range filenames {
// 		fmt.Printf("Criando Matriz de Adjacências\n")

// 		A, err := CreateMatrix(name)
// 		if err != nil {
// 			t.Fatalf("Error %v in reading file.", err)
// 		}

// 		fmt.Printf("Matriz Criada\n")

// 		fmt.Printf("Arquivo de entrada: %v\n", name)

// 		for _, v := range testing_vertices {
// 			_, parent, _ := A.BFS(v)
// 			fmt.Printf("Raiz: %v - Vértice: %v - Pai: %v\n", v, 10, parent[9])
// 			fmt.Printf("Raiz: %v - Vértice: %v - Pai: %v\n", v, 20, parent[19])
// 			fmt.Printf("Raiz: %v - Vértice: %v - Pai: %v\n", v, 30, parent[29])
// 		}

// 		fmt.Println()
// 	}
// }

// func TestDistances(t *testing.T) {
// 	testing_vertices := [3]int{10, 20, 30}

// 	for _, name := range filenames {
// 		fmt.Printf("Criando Matriz de Adjacências\n")

// 		A, err := CreateMatrix(name)
// 		if err != nil {
// 			t.Fatalf("Error %v in reading file.", err)
// 		}

// 		fmt.Printf("Matriz Criada\n")

// 		fmt.Printf("Arquivo de entrada: %v\n", name)
// 		for i := 0; i < 3; i++ {
// 			for j := i + 1; j < 3; j++ {
// 				dist := A.Distance(testing_vertices[i], testing_vertices[j])
// 				fmt.Printf("(%v, %v) - %v\n", testing_vertices[i], testing_vertices[j], dist)
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func TestConnectedComponents(t *testing.T) {
// 	for _, name := range filenames {
// 		fmt.Printf("Criando Matriz de Adjacências\n")

// 		A, err := CreateMatrix(name)
// 		if err != nil {
// 			t.Fatalf("Error %v in reading file.", err)
// 		}

// 		fmt.Printf("Matriz Criada\n")

// 		components, graphs := A.ConectedComponents()

// 		fmt.Printf("Grafo de %v\nComponentes Conexas: %v\nMaior Componente: %v\nMenor Componente: %v\n\n",
// 			name, graphs, components[graphs-1].Size, components[0].Size)
// 	}
// }

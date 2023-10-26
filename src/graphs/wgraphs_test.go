package graphs

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Corretor, edite estas variáveis para que possa ler os arquivos corretos!
var filenamesw = []string{
	// "../../../grafos/TP2/grafo_W_1.txt",
	// "../../../grafos/TP2/grafo_W_2.txt",
	// "../../../grafos/TP2/grafo_W_3.txt",
	"../../../grafos/TP2/grafo_W_4.txt",
	// "../../../grafos/TP2/grafo_W_5.txt",
}

var filenamerede = []string{
	"../../../grafos/TP2/rede_colaboracao.txt",
	"../../../grafos/TP2/rede_colaboracao_vertices.txt",
}

func TestDistancesDKT(t *testing.T) {
	vertices := []int{20, 30, 40, 50, 60}
	for _, name := range filenamesw {
		A := WList()

		ReadFile(A, name)
		fmt.Printf("Grafo Lido\n")

		t1 := time.Now()

		dists, parent := DijkstraHeap(A, 10)

		t2 := time.Now()

		t := t2.Sub(t1).Seconds()

		fmt.Printf("DKT Feito para arquivo %v\nUsou: Heap\nLevou %v\n", name, t)

		for _, v := range vertices {
			way := MyWay(parent, v)
			fmt.Printf("%v - Custo: %0.2f\n\n", way, dists[v-1])
		}
	}
}

func TestTimesDKTHeap(t *testing.T) {
	I := 10

	for _, name := range filenamesw {
		A := WList()

		ReadFile(A, name)
		fmt.Printf("Grafo Lido\n")

		var t float64 = 0

		for i := 0; i < I; i++ {
			t1 := time.Now()

			v := rand.Intn(A.Vertices)

			DijkstraHeap(A, v)

			t2 := time.Now()

			t += t2.Sub(t1).Seconds()
		}

		fmt.Printf("DKT Feito para arquivo %v\nUsou: Heap\nMedia: %v\n\n", name, t/float64(I))
	}
}

func TestTimesDKTVector(t *testing.T) {
	I := 1

	for _, name := range filenamesw {
		A := WList()

		ReadFile(A, name)
		fmt.Printf("Grafo Lido\n")

		var t float64 = 0

		for i := 0; i < I; i++ {
			t1 := time.Now()

			Dijkstra(A, 10)

			t2 := time.Now()

			t += t2.Sub(t1).Seconds()
		}

		fmt.Printf("DKT Feito para arquivo %v\nUsou: Vetor\nMedia: %v\n", name, t/float64(I))
	}
}

func TestDistancesDKTRede(t *testing.T) {

	mapVertexName, mapNameVertex := ReadFileNet(filenamerede[1])

	start := mapNameVertex["Edsger W. Dijkstra"]
	// Problema com nome da pequena Eva!!!
	eva_number := 11386
	names := []string{"Alan M. Turing", "J. B. Kruskal", "Jon M. Kleinberg", mapVertexName[eva_number], "Daniel R. Figueiredo"}
	names_vertex := []int{0, 0, 0, 0, 0}
	for i, name := range names {
		names_vertex[i] = mapNameVertex[name]
	}

	A := WList()

	ReadFile(A, filenamerede[0])
	fmt.Printf("Grafo Lido\n")

	t1 := time.Now()

	fmt.Printf("Start: %v\n", start)
	dists, parent := DijkstraHeap(A, start)

	t2 := time.Now()

	time := t2.Sub(t1).Seconds()

	fmt.Printf("DKT Feito para arquivo %v\nUsou: Heap\nLevou %v\n", filenamerede[0], time)

	for _, v := range names_vertex {
		way := MyWay(parent, v)
		for _, vertex := range way {
			fmt.Printf(" %v -", mapVertexName[vertex])
		}
		fmt.Printf("Custo: %0.2f\n\n", dists[v-1])
	}
}

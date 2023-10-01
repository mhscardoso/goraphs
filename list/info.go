package list

import (
	"fmt"
	"os"

	"github.com/mhscardoso/goraphs/queue"
	"github.com/mhscardoso/goraphs/sort"
)

func (L *List) GetInfo() (int, int, int, int, float32, float32) {
	edges := make([]int, L.N)
	var sum float32 = 0
	for i, v := range L.Vector {
		edges[i] = GetDegree(v)
		sum += float32(edges[i])
	}

	median := sum / float32(L.N)

	sort.Sort(&edges)

	var middle float32

	if L.N%2 == 0 {
		middle = (float32(edges[L.N/2]) + float32(edges[(L.N/2)-1])) / 2
	} else {
		middle = float32(edges[L.N/2]) / 2
	}

	return L.N, L.M, edges[0], edges[L.N-1], median, middle
}

func GetDegree(e Element) int {
	sum := 0
	for n := e.Next; n != nil; n = n.Next {
		sum++
	}
	return sum
}

func (L *List) SaveData(filename string) {
	N_a, M_a, min_a, max_a, median_a, middle_a := L.GetInfo()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	str := fmt.Sprintf("Vertices: %v\nEdges: %v\nMin Degree: %v\nMax Degree: %v\nMean Degrees: %v\nMedian Degrees: %v\n",
		N_a, M_a, min_a, max_a, median_a, middle_a)

	_, err2 := f.WriteString(str)
	if err2 != nil {
		panic(err2)
	}

	// Data about conected components
	components := make([]*queue.Queue, L.N)
	component := 0
	signal := make([]byte, L.N)

	for i := 0; i < L.N; i++ {
		if signal[i] == 1 {
			continue
		}
		component++
		components[component-1] = queue.New()
		L.BFS_with_known_components(
			L.Vector[i].Vertex+1,
			&signal,
			components[component-1])
	}

	for i := 0; i < component; i++ {
		if components[i].Size > 1 {
			fmt.Printf("C: %v -- S: %v\n", i, components[i].Size)
		}
	}
}

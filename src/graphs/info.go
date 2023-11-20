package graphs

import (
	"fmt"
	"os"

	"github.com/mhscardoso/goraphs/container/queue"
	"github.com/mhscardoso/goraphs/container/set"
	//"github.com/mhscardoso/goraphs/sort"
)

func GetInfo(g Graph) (vertices int, edges int, lowest int, greatest int, median float32, middle float32) {
	vertices = g.N()
	edges = g.M()

	edgesVector := make([]int, vertices)
	sum := 0
	for i := 0; i < vertices; i++ {
		edgesVector[i] = GetDegree(g, i)
		sum += edgesVector[i]
	}

	median = float32(sum) / float32(vertices)

	fmt.Printf("Vertices: %v\nEdgesx2: %v\nSum: %v\nValidate: %v\n",
		vertices, 2*edges, sum, sum == 2*edges)

	//sort.Sort(&edgesVector)

	if vertices%2 == 0 {
		middle = (float32(edgesVector[vertices/2]) + float32(edgesVector[(vertices/2)-1])) / 2
	} else {
		middle = float32(edgesVector[vertices/2]) / 2
	}

	lowest = edgesVector[0]
	greatest = edgesVector[vertices-1]

	return
}

func GetDegree(g Graph, vertex int) int {
	var n1 set.Set[int]
	var n2 set.SetW[int, float64]

	n1, ok := g.Neighbors(vertex).(set.Set[int])
	if !ok || n1 == nil {
		n2 = g.Neighbors(vertex).(set.SetW[int, float64])
		return len(n2)
	}

	return len(n1)
}

func ConectedComponents(g Graph) *queue.Queue[*queue.Queue[int]] {
	vertices := g.N()

	components := new(queue.Queue[*queue.Queue[int]])
	signalAll := make([]byte, vertices)
	signals := make([]byte, vertices)

	for i := 0; i < vertices; i++ {
		if signalAll[i] == 1 {
			continue
		}

		fmt.Printf("Vertex: %v\n", i+1)

		components.Insert(new(queue.Queue[int])) // Add in Last

		BFS(g, i+1, 0, &signals)

		for p := range signals {
			if signals[p] == 1 {
				components.Last.Vertex.Insert(p + 1)
				signalAll[p] = 1
			}
			signals[p] = 0
		}
	}

	return components
}

func WriteConnected(filename string, results *queue.Queue[*queue.Queue[int]]) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}

	defer f.Close()

	for e := results.First; e != nil; e = e.Next {
		f.WriteString(fmt.Sprintf("%v\n", e.Vertex.Length))

		for q := e.Vertex.First; q != nil; q = q.Next {
			f.WriteString(fmt.Sprintf("%v\n", q.Vertex))
		}
		f.WriteString("---------------------------------------------------------------------------------\n")
	}
}

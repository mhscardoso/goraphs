package graphs

import (
	"github.com/mhscardoso/goraphs/queue"
	"github.com/mhscardoso/goraphs/sort"
)

func GetInfo(g Graph) (int, int, int, int, float32, float32, []int) {
	vertices := g.N()
	edges := g.M()

	edgesVector := make([]int, vertices)
	var sum float32 = 0
	for i := 0; i < vertices; i++ {
		edgesVector[i] = GetDegree(g, i)
		sum += float32(edgesVector[i])
	}

	median := sum / float32(vertices)

	sort.Sort(&edgesVector)

	var middle float32

	if vertices%2 == 0 {
		middle = (float32(edgesVector[vertices/2]) + float32(edgesVector[(vertices/2)-1])) / 2
	} else {
		middle = float32(edgesVector[vertices/2]) / 2
	}

	return vertices, edges, edgesVector[0], edgesVector[vertices-1], median, middle, edgesVector
}

func GetDegree(g Graph, vertex int) int {
	neighbors := g.Neighbors(vertex)

	degree := 0

	for e := neighbors; e != nil; e = e.Next {
		degree++
	}

	return degree
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

		components.Insert(new(queue.Queue[int])) // Add in Last

		BFS(g, i+1, &signals)

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

package graphs

import (
	"github.com/mhscardoso/goraphs/queue"
	"github.com/mhscardoso/goraphs/sort"
)

func GetInfo(g Graph) (int, int, int, int, float32, float32) {
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

	return vertices, edges, edgesVector[0], edgesVector[vertices-1], median, middle
}

func GetDegree(g Graph, vertex int) int {
	neighbors := g.Neighbors(vertex)

	degree := 0

	for e := neighbors; e != nil; e = e.Next {
		degree++
	}

	return degree
}

func ConectedComponents(g Graph) ([]*queue.Queue, int) {
	vertices := g.N()

	components := make([]*queue.Queue, vertices)
	component := 0
	signal := make([]byte, vertices)

	for i := 0; i < vertices; i++ {
		if signal[i] == 1 {
			continue
		}

		component++
		_, _, components[component-1] = BFS(g, i+1, &signal)
	}

	return components[:component], component
}

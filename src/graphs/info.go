package graphs

import (
	"github.com/mhscardoso/goraphs/queue"
	"github.com/mhscardoso/goraphs/sort"
)

func GetInfo(g Graph) (vertices int, edges int, lowest int, greatest int, median float32, middle float32) {
	vertices = g.N()
	edges = g.M()

	edgesVector := make([]int, vertices)
	var sum float32 = 0
	for i := 0; i < vertices; i++ {
		edgesVector[i] = GetDegree(g, i)
		sum += float32(edgesVector[i])
	}

	median = sum / float32(vertices)

	sort.Sort(&edgesVector)

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

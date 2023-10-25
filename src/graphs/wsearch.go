package graphs

import (
	//"fmt"
	"fmt"
	"math"

	"github.com/mhscardoso/goraphs/container/set"
)

func MinDist(S set.Set[int], distances []float64) (int, float64) {
	var min float64 = math.Inf(1)
	var u int

	for k := range S {
		if distances[k] < min {
			u = k
			min = distances[k]
		}
	}

	return u, min
}

func Dijkstra(g Graph, s int) []float64 {
	vertex := s - 1

	if s <= 0 || s > g.N() {
		panic("S must be in the interval (1, vertices)\n")
	}

	// Actually, this is V-S Set
	S := make(set.Set[int])
	for i := 0; i < g.N(); i++ {
		S.Add(i)
	}

	distances := make([]float64, g.N())
	for i := range distances {
		distances[i] = math.Inf(1)
	}

	distances[vertex] = 0

	neighbors, ok := g.Neighbors(vertex).(set.SetW[int])
	if !ok {
		return nil
	}

	for len(S) != 0 {
		u, _ := MinDist(S, distances)
		S.Remove(u)

		neighbors = g.Neighbors(u).(set.SetW[int])

		for k, v := range neighbors {
			if v < 0 {
				fmt.Printf("Negative weigths are not supported\n")
				return nil
			}
			if distances[k] > distances[u]+v {
				distances[k] = distances[u] + v
			}
		}
	}

	return distances
}

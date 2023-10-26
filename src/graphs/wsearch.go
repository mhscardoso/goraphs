package graphs

import (
	"fmt"
	"math"
	"os"

	"github.com/mhscardoso/goraphs/container/heap"
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

func Dijkstra(g Graph, s int) ([]float64, []int) {
	vertex := s - 1

	if s <= 0 || s > g.N() {
		panic("S must be in the interval (1, vertices)\n")
	}

	parents := make([]int, g.N())

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
		fmt.Printf("Only graphs with weigths supported\n")
		return nil, nil
	}

	for len(S) != 0 {
		u, _ := MinDist(S, distances)
		S.Remove(u)

		neighbors = g.Neighbors(u).(set.SetW[int])

		for k, v := range neighbors {
			if v < 0 {
				fmt.Printf("Negative weigths are not supported\n")
				return nil, nil
			}

			if distances[k] > distances[u]+v {
				distances[k] = distances[u] + v
				parents[k] = u + 1
			}
		}
	}

	return distances, parents
}

func DijkstraHeap(g Graph, s int) (dist []float64, parents []int) {
	vertex := s - 1

	if s <= 0 || s > g.N() {
		panic("S must be in the interval (1, vertices)\n")
	}

	parents = make([]int, g.N())

	dist = make([]float64, g.N())

	// Actually, this is V-S Set
	S := make(set.Set[int])
	for i := 0; i < g.N(); i++ {
		S.Add(i)
	}

	distances := heap.New(g.N())
	distances.Update(vertex, 0)

	neighbors, ok := g.Neighbors(vertex).(set.SetW[int])
	if !ok {
		fmt.Printf("Only graphs with weigths supported\n")
		return nil, nil
	}

	for len(S) != 0 {
		u := distances.Remove()
		du := u.Distance
		S.Remove(u.Vertex)

		neighbors = g.Neighbors(u.Vertex).(set.SetW[int])

		for k, v := range neighbors {
			if v < 0 {
				fmt.Printf("Negative weigths are not supported\n")
				return nil, nil
			}

			if distances.Position[k] == -1 {
				continue
			}

			if distances.Vector[distances.Position[k]].Distance > du+v {
				distances.Update(k, du+v)
				dist[k] = du + v
				parents[k] = u.Vertex + 1
			}
		}
	}

	return
}

func WriteResults(filename string, results []float64, parents []int, root int) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}

	defer f.Close()

	f.WriteString(fmt.Sprintf("%v\n", root))

	for i := range results {
		f.WriteString(fmt.Sprintf("%0.2f - %v - %v\n", results[i], i+1, parents[i]))
	}
}

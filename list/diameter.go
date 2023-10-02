package list

import (
	"math/rand"
)

// All algorithms to the shortest path

// func (L *List) BFSrn() {
// 	rn := int(math.Sqrt(float64(L.N)))

// }

func (L *List) DiameterApprox() int {
	first_vertex := rand.Intn(L.N)

	_, _, level := L.BFS(first_vertex)

	max_level := 0

	for i := range level {
		if level[i] > level[max_level] {
			max_level = i
		}
	}

	_, _, level2 := L.BFS(max_level + 1)

	max_level = 0

	for i := range level {
		if level2[i] > level2[max_level] {
			max_level = i
		}
	}

	return int(level2[max_level])
}

func (L *List) Diameter() uint {
	var max_distance uint = 0

	for i := range L.Vector {
		for j := i + 1; j < L.N; j++ {
			dist := L.Distance(i+1, j)

			if dist > int(max_distance) {
				max_distance = uint(dist)
			}
		}
	}

	return max_distance
}

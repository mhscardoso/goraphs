package list

import (
	"math/rand"
)

// All algorithms to the shortest path

// func (L *List) BFSrn() {
// 	rn := int(math.Sqrt(float64(L.N)))

// }

func (L *List) DiameterApprox() int {

	max_level := 0

	for i := 0; i < 1000; i++ {

		first_vertex := rand.Intn(L.N)

		_, _, level := L.BFS(first_vertex)

		for _, v := range level {
			if v > uint(max_level) {
				max_level = int(v)
			}
		}

	}

	return max_level
}

func (L *List) Diameter() uint {
	var max_distance uint = 0

	for i := range L.Vector {
		for j := i + 2; j < L.N; j++ {
			dist := L.Distance(i+1, j)

			if dist > int(max_distance) {
				max_distance = uint(dist)
			}
		}
	}

	return max_distance
}

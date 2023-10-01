package matrix

import (
	"github.com/mhscardoso/goraphs/sort"
)

func (A *AdjMatrix) GetInfo() (int, int, int, int, float32, float32) {
	edges := make([]int, A.N)
	var sum float32 = 0
	for i, v := range A.P {
		edges[i] = GetDegree(v)
		sum += float32(edges[i])
	}

	median := sum / float32(A.N)

	sort.Sort(&edges)

	var middle float32

	if A.N%2 == 0 {
		middle = (float32(edges[A.N/2]) + float32(edges[(A.N/2)-1])) / 2
	} else {
		middle = float32(edges[A.N/2]) / 2
	}

	return A.N, A.M, edges[0], edges[A.N-1], median, middle
}

// Get the degree of a given vertex in matrix
func GetDegree(vertex []byte) int {
	N := len(vertex)

	neighbors := 0

	for j := 0; j < N; j++ {
		supposed := vertex[j]
		if supposed != 0 {
			for bit := 0; bit < 8; bit++ {
				if supposed&byte(1<<(7-bit)) != 0 {
					neighbors++
				}
			}
		}
	}
	return neighbors
}

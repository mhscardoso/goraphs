package matrix

import (
	"fmt"
)

func (A *AdjMatrix) GetInfo() (int, int, int, int, float32, float32, []int) {
	edges := make([]int, A.N)
	var sum float32 = 0
	for i := range A.P {
		edges[i] = GetDegree(A.P[i])
		sum += float32(edges[i])
	}

	fmt.Printf("Sum Matrix: %v\nSum Matrix: %v\nA.M: %v\n\n", sum == float32(2*A.M), sum, 2*A.M)

	median := sum / float32(A.N)

	//sort.Sort(&edges)

	var middle float32

	if A.N%2 == 0 {
		middle = (float32(edges[A.N/2]) + float32(edges[(A.N/2)-1])) / 2
	} else {
		middle = float32(edges[A.N/2]) / 2
	}

	return A.N, A.M, edges[0], edges[A.N-1], median, middle, edges
}

// Get the degree of a given vertex in matrix
func GetDegree(vertex []byte) int {
	neighbors := 0

	for _, v := range vertex {
		if v != 0 {
			for bit := 0; bit < 8; bit++ {
				if v&byte(0b00000001<<(7-bit)) != 0 {
					neighbors++
				}
			}
		}
	}
	return neighbors
}

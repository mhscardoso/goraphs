package list

import (
	"fmt"
)

func (L *List) GetInfo() (int, int, int, int, float32, float32, []int) {
	edges := make([]int, L.N)
	var sum float32 = 0
	for i, v := range L.Vector {
		edges[i] = GetDegree(v)
		sum += float32(edges[i])
	}

	fmt.Printf("Sum List: %v\nSum List %v\nL.M: %v\n\n", sum == float32(2*L.M), sum, 2*L.M)

	median := sum / float32(L.N)

	//sort.Sort(&edges)

	var middle float32

	if L.N%2 == 0 {
		middle = (float32(edges[L.N/2]) + float32(edges[(L.N/2)-1])) / 2
	} else {
		middle = float32(edges[L.N/2]) / 2
	}

	return L.N, L.M, edges[0], edges[L.N-1], median, middle, edges
}

func GetDegree(e Element) int {
	sum := 0
	for n := e.Next; n != nil; n = n.Next {
		sum++
	}
	return sum
}

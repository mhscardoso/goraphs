package matrix

import (
	"fmt"
	"os"

	"github.com/mhscardoso/goraphs/sort"
)

func (A *AdjMatrix) GetInfo() (int, int, int, int, float32, float32) {
	edges := make([]int, A.N)
	var sum float32 = 0
	for i := range A.P {
		edges[i] = GetDegree(A.P[i])
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

func (A *AdjMatrix) SaveData(filename string) {
	N_a, M_a, min_a, max_a, median_a, middle_a := A.GetInfo()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	str := fmt.Sprintf("Vertices: %v\nEdges: %v\nMin Degree: %v\nMax Degree: %v\nMean Degrees: %v\nMedian Degrees: %v\n",
		N_a, M_a, min_a, max_a, median_a, middle_a)

	_, err2 := f.WriteString(str)
	if err2 != nil {
		panic(err2)
	}
}

package matrix

import (
	"math/rand"
	"time"
)

func (A *AdjMatrix) MeasureBFS(s int) float64 {
	t1 := time.Now()

	A.BFS(s)

	t2 := time.Now()

	diff := t2.Sub(t1)

	return float64(diff.Seconds())
}

func (A *AdjMatrix) StatsBFS() float64 {
	var time float64 = 0
	for i := 1; i <= 100; i++ {
		s := rand.Intn(A.N)
		time += float64(A.MeasureBFS(s))
	}

	return time / float64(100)
}

func (A *AdjMatrix) MeasureDFS(s int) float64 {
	t1 := time.Now()

	A.DFS(s)

	t2 := time.Now()

	diff := t2.Sub(t1)

	return float64(diff.Seconds())
}

func (A *AdjMatrix) StatsDFS() float64 {
	var time float64 = 0
	for i := 1; i <= 100; i++ {
		time += float64(A.MeasureDFS(i))
	}

	return time / float64(100)
}

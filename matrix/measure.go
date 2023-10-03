package matrix

import (
	"time"
)

func (A *AdjMatrix) MeasureBFS(s int) float64 {
	t1 := time.Now()

	A.BFS(s)

	t2 := time.Now()

	diff := t2.Sub(t1)

	return float64(diff.Seconds())
}

func (A *AdjMatrix) MeasureDFS(s int) float64 {
	t1 := time.Now()

	A.DFS(s)

	t2 := time.Now()

	diff := t2.Sub(t1)

	return float64(diff.Seconds())
}

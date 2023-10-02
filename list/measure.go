package list

import (
	"math/rand"
	"time"
)

func (L *List) MeasureBFS(s int) float64 {
	t1 := time.Now()

	L.BFS(s)

	t2 := time.Now()

	diff := t2.Sub(t1)

	return float64(diff.Seconds())
}

func (L *List) StatsBFS() float64 {
	var time float64 = 0
	for i := 1; i <= 100; i++ {
		s := rand.Intn(L.N)
		time += float64(L.MeasureBFS(s))
	}

	return time / float64(100)
}

func (L *List) MeasureDFS(s int) float64 {
	t1 := time.Now()

	L.DFS(s)

	t2 := time.Now()

	diff := t2.Sub(t1)

	return float64(diff.Seconds())
}

func (L *List) StatsDFS() float64 {
	var time float64 = 0
	for i := 1; i <= 100; i++ {
		time += float64(L.MeasureDFS(i))
	}

	return time / float64(100)
}

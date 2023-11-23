package graphs

import (
	"fmt"
	"os"
	"time"

	"github.com/mhscardoso/goraphs/container/queue"
	"github.com/mhscardoso/goraphs/container/set"
	"github.com/mhscardoso/goraphs/src/awlists"
	"github.com/mhscardoso/goraphs/src/flist"
)

func FordFulkerson(l *flist.FList, s int, t int, delta int) {
	initial_delta := delta

	residual := ResidualGraph(l, initial_delta)

	for {
		P := MyWay(BFSWeigth(residual, s, t, nil), t)

		for len(P) == 1 && initial_delta >= 1 {
			initial_delta /= 2
			if initial_delta == 0 {
				return
			}
			AugmentResidual(l, residual, initial_delta)
			P = MyWay(BFSWeigth(residual, s, t, nil), t)
		}

		Augment(l, residual, P)
		UpdateResidual(l, residual, P, initial_delta)
	}
}

// Get residual graph given other
func ResidualGraph(l *flist.FList, delta int) *awlists.WList[int] {
	// Residual graph is targeted
	residual := new(awlists.WList[int])
	residual.Targeted = true
	residual.Edges = 0

	residual.Allocate(l.Vertices)

	// Iterating over all vertices
	for i, w := range l.Vector {
		for n, v := range w {
			if v.GetWeigth()-v.GetFlow() >= delta {
				residual.Relate(i+1, int(n)+1, v.GetWeigth()-v.GetFlow(), &residual.Edges)
			}
		}
	}

	return residual
}

func UpdateResidual(l *flist.FList, res *awlists.WList[int], P []int, delta int) {
	for i := 1; i < len(P); i++ {
		if l.Vector[P[i]].Contains(uint32(P[i-1])) {
			res.Vector[P[i]][P[i-1]] = l.Vector[P[i]][uint32(P[i-1])].GetDiff()
			res.Vector[P[i-1]][P[i]] = l.Vector[P[i]][uint32(P[i-1])].GetFlow()
		} else {
			res.Vector[P[i-1]][P[i]] = l.Vector[P[i-1]][uint32(P[i])].GetDiff()
			res.Vector[P[i]][P[i-1]] = l.Vector[P[i-1]][uint32(P[i])].GetFlow()
		}

		// Deleting edges
		if res.Vector[P[i]][P[i-1]] < delta {
			res.Vector[P[i]].Remove(P[i-1])
		}

		if res.Vector[P[i-1]][P[i]] < delta {
			res.Vector[P[i-1]].Remove(P[i])
		}
	}
}

func AugmentResidual(l *flist.FList, res *awlists.WList[int], delta int) {
	// Iterating over all vertices
	for i, w := range l.Vector {
		for n, v := range w {
			if v.GetWeigth()-v.GetFlow() >= delta {
				res.Vector[i][int(n)] = v.GetWeigth() - v.GetFlow()
			} else {
				res.Vector[i].Remove(int(n))
			}

			if v.GetFlow() >= delta {
				res.Vector[int(n)][i] = v.GetFlow()
			} else {
				res.Vector[n].Remove(i)
			}
		}
	}
}

func Bottleneck(res *awlists.WList[int], P []int) int {
	if len(P) == 1 {
		return 0
	}

	b := res.Vector[P[1]][P[0]]

	for i := 2; i < len(P); i++ {
		b = min(b, res.Vector[P[i]][P[i-1]])
	}

	return b
}

func Augment(l *flist.FList, res *awlists.WList[int], P []int) int {
	b := Bottleneck(res, P)

	if b == 0 {
		return b
	}

	for i := 1; i < len(P); i++ {
		if l.Vector[P[i]].Contains(uint32(P[i-1])) {
			actualFlow := l.Vector[P[i]][uint32(P[i-1])].GetFlow()
			l.Vector[P[i]].UpdateFlow(uint32(P[i-1]), actualFlow+b)
		} else {
			actualFlow := l.Vector[P[i-1]][uint32(P[i])].GetFlow()
			l.Vector[P[i]].UpdateFlow(uint32(P[i]), actualFlow-b)
		}
	}

	return b
}

/* New BFS written to support
 * graphs with integer weigths
 * -- Calculates a path in the
 * -- residual graph.
 */
func BFSWeigth(l *awlists.WList[int], s int, d int, known_signals *[]byte) (parent []int) {
	vertices := l.Vertices

	if known_signals != nil && len(*known_signals) != vertices {
		panic("The Known Signals length must have the same number of vertices\n")
	}

	if s == d || s <= 0 || s > vertices || d < 0 || d > vertices {
		panic("S must be in the interval (1, vertices)\nS and D must be differents\nD must be in the interval (0, vertices)")
	}

	var signal *[]byte

	// Signal to mark a vertex when discovered
	if known_signals == nil {
		signals := make([]byte, vertices)
		signal = &signals
	} else {
		signal = known_signals
	}

	// parent[i] = j == vertex j is parent of i in BFS tree
	parent = make([]int, vertices)

	// Dealing with vertices between 0 -- N-1
	vertex := s - 1

	// Start a Queue
	Q := queue.New[int]()

	// Mark given vertex and insert in queue
	(*signal)[vertex] = 1
	Q.Insert(vertex)

	// If the graph has weigths it fails here
	neighbors, ok := l.Neighbors(vertex).(set.SetW[int, int])
	if !ok {
		return
	}

	// For each vertex in queue; While queue not empty
	// It remeves a vertex in each iteration
	for e := Q.Remove(); e != nil; e = Q.Remove() {

		// Vertex being explored and its level in tree
		exploring := e.Vertex

		neighbors = l.Neighbors(exploring).(set.SetW[int, int])

		// For each neighbor of vertex
		for w := range neighbors {
			if (*signal)[w] == 0 {
				(*signal)[w] = 1          // Mark vertex
				Q.Insert(w)               // Insert in queue
				parent[w] = exploring + 1 // Save its parent

				if w+1 == d {
					return
				}
			}
		}
	}

	return
}

func WriteFordResult(l *flist.FList, s int, t int, delta int, filename string) {
	fmt.Printf("Realizando o Ford-Fulkerson\n")
	FordFulkerson(l, s, t, delta)

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	maxFlow := 0

	for _, v := range l.Vector[s-1] {
		maxFlow += v.GetFlow()
	}

	fmt.Printf("Escrevendo o arquivo em %v\n", filename)
	f.WriteString(fmt.Sprintf("Máximo entre %v e %v: %v\n", s, t, maxFlow))

	for i, v := range l.Vector {
		for k, w := range v {
			f.WriteString(fmt.Sprintf("%v - %v - %v\n", i+1, k+1, w.GetFlow()))
		}
	}

	fmt.Printf("Arquivo escrito.\n")
}

func FordFulkersonStats(l *flist.FList, s int, t int, delta int) (int, float64, float64, int, float64) {
	initial_delta := delta

	t1 := time.Now()
	residual := ResidualGraph(l, initial_delta)
	t2 := time.Now()

	residualTime := t2.Sub(t1).Seconds()

	BFStimes := 0
	var BFStotal float64 = 0
	augs := 0
	var AugTime float64 = 0

	for {
		BFStimes++
		t1 = time.Now()
		P := MyWay(BFSWeigth(residual, s, t, nil), t)
		t2 = time.Now()

		BFStotal += t2.Sub(t1).Seconds()

		for len(P) == 1 && initial_delta >= 1 {
			initial_delta /= 2
			if initial_delta == 0 {
				return BFStimes, BFStotal, residualTime, augs, AugTime
			}

			augs++
			t1 = time.Now()
			AugmentResidual(l, residual, initial_delta)
			t2 = time.Now()

			AugTime += t2.Sub(t1).Seconds()

			BFStimes++
			t1 = time.Now()
			P = MyWay(BFSWeigth(residual, s, t, nil), t)
			t2 = time.Now()

			BFStotal += t2.Sub(t1).Seconds()
		}

		Augment(l, residual, P)
		UpdateResidual(l, residual, P, initial_delta)
	}
}

package graphs

import (
	"github.com/mhscardoso/goraphs/container/queue"
	"github.com/mhscardoso/goraphs/container/set"
	"github.com/mhscardoso/goraphs/src/awlists"
	"github.com/mhscardoso/goraphs/src/flist"
)

func FordFulkerson(l *flist.FList, s int, t int) {

}

// Get residual graph given other
func ResidualGraph(l *flist.FList) *awlists.WList[int] {
	// Residual graph is targeted
	residual := new(awlists.WList[int])
	residual.Targeted = true
	residual.Edges = 0

	// Iterating over all vertices
	for i := range l.Vector {
		for n, v := range l.Vector[i] {
			residual.Relate(i, int(n), v.GetWeigth(), &residual.Edges)
		}
	}

	return residual
}

/* New BFS written to support
 * graphs with flows.
 */
func BFSFlow(l *flist.FList, s int, d int, known_signals *[]byte) (parent []int, level []uint) {
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

	// Level of vertices in tree of BFS
	level = make([]uint, vertices)

	// Dealing with vertices between 0 -- N-1
	vertex := s - 1

	// Start a Queue
	Q := queue.New[int]()

	// Mark given vertex and insert in queue
	(*signal)[vertex] = 1
	Q.Insert(vertex)

	// If the graph has weigths it fails here
	neighbors, ok := l.Neighbors(vertex).(set.SetF)
	if !ok {
		return
	}

	// For each vertex in queue; While queue not empty
	// It remeves a vertex in each iteration
	for e := Q.Remove(); e != nil; e = Q.Remove() {

		// Vertex being explored and its level in tree
		exploring := e.Vertex
		level_exploring_plus := level[exploring] + 1

		neighbors = l.Neighbors(exploring).(set.SetF)

		// For each neighbor of vertex
		for w := range neighbors {
			if (*signal)[w] == 0 {
				(*signal)[w] = 1                // Mark vertex
				Q.Insert(int(w))                // Insert in queue
				parent[w] = exploring + 1       // Save its parent
				level[w] = level_exploring_plus // Save its level

				if int(w+1) == d {
					return
				}
			}
		}
	}

	return
}

func Bottleneck(l *flist.FList, s int, t int) int {
	parents, _ := BFSFlow(l, s, t, nil)

	ls := MyWay(parents, t)

	b := l.Vector[ls[1]-1][uint32(ls[0]-1)][0]

	for i := 2; i < len(ls); i++ {
		b = min(b, l.Vector[ls[i]-1][uint32(ls[i-1]-1)][0])
	}

	return b
}

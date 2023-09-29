package list

import "github.com/mhscardoso/goraphs/queue"

func (L *List) BFS(s int) ([]int, []uint) {

	if s <= 0 || s > L.N {
		return nil, nil
	}

	// Signal to mark a vertex when discovered
	signal := make([]byte, L.N)

	// parent[i] = j == vertex j is parent of i in BFS tree
	parent := make([]int, L.N)

	// Level of vertices in tree of BFS
	level := make([]uint, L.N)

	// Dealing with vertices between 0 -- N-1
	vertex := s - 1

	// Start a Queue
	Q := queue.New()

	// Mark given vertex and insert in queue
	signal[vertex] = 1
	Q.Insert(vertex)

	// For each vertex in queue; While queue not empty
	// It remeves a vertex in each iteration
	for e := Q.Remove(); e != nil; e = Q.Remove() {

		// Vertex being explored and its level in tree
		exploring := e.Vertex
		level_exploring_plus := level[exploring] + 1

		// For each neighbor of vertex
		for w := L.Vector[exploring].Next; w != nil; w = w.Next {
			if signal[w.Vertex] == 0 {
				signal[w.Vertex] = 1                   // Mark vertex
				Q.Insert(w.Vertex)                     // Insert in queue
				parent[w.Vertex] = exploring + 1       // Save its parent
				level[w.Vertex] = level_exploring_plus // Save its level
			}
		}
	}

	return parent, level
}

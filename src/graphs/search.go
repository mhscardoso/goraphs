package graphs

import (
	"github.com/mhscardoso/goraphs/queue"
	"github.com/mhscardoso/goraphs/stack"
)

// Executes the BFS
func BFS(g Graph, s int, known_signals *[]byte) ([]int, []uint) {
	vertices := g.N()

	if len(*known_signals) != vertices {
		panic("The Known Signals length must have the same number of vertices\n")
	}

	if s <= 0 || s > vertices {
		panic("S must be in the interval (1, vertices)\n")
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
	parent := make([]int, vertices)

	// Level of vertices in tree of BFS
	level := make([]uint, vertices)

	// Dealing with vertices between 0 -- N-1
	vertex := s - 1

	// Start a Queue
	Q := queue.New[int]()

	// Mark given vertex and insert in queue
	(*signal)[vertex] = 1
	Q.Insert(vertex)

	// For each vertex in queue; While queue not empty
	// It remeves a vertex in each iteration
	for e := Q.Remove(); e != nil; e = Q.Remove() {

		// Vertex being explored and its level in tree
		exploring := e.Vertex
		level_exploring_plus := level[exploring] + 1

		neighbors := g.Neighbors(exploring)

		// For each neighbor of vertex
		for w := neighbors; w != nil; w = w.Next {
			if (*signal)[w.Vertex] == 0 {
				(*signal)[w.Vertex] = 1                // Mark vertex
				Q.Insert(w.Vertex)                     // Insert in queue
				parent[w.Vertex] = exploring + 1       // Save its parent
				level[w.Vertex] = level_exploring_plus // Save its level
			}
		}
	}

	return parent, level
}

// Executes the DFS
func DFS(g Graph, s int, known_signals *[]byte) ([]int, []uint) {
	vertices := g.N()

	if len(*known_signals) != vertices {
		panic("The Known Signals length must have the same number of vertices\n")
	}

	if s <= 0 || s > vertices {
		panic("S must be in the interval (1, vertices)\n")
	}

	// Working from 0 to N-1
	vertex := s - 1

	// To mark a vertex
	// Signal to mark a vertex when discovered
	signal := known_signals
	if signal == nil {
		signals := make([]byte, vertices)
		signal = &signals
	}

	// Parent os vertices in DFS
	parent := make([]int, vertices)

	// Level of vertices in DFS
	level := make([]uint, vertices)

	// Start a new stack and inserts the vertex in it
	P := stack.New[int]()
	P.Insert(vertex)

	// While P is not empty
	for e := P.Remove(); e != nil; e = P.Remove() {
		exploring := e.Vertex
		level_plus_one := level[exploring] + 1

		if (*signal)[exploring] == 0 {
			(*signal)[exploring] = 1

			neighbors := g.Neighbors(exploring)

			for w := neighbors; w != nil; w = w.Next {
				P.Insert(w.Vertex) // Inserts neighbor in stack

				if w.Vertex != vertex && parent[w.Vertex] == 0 {
					parent[w.Vertex] = exploring + 1 // Save parent
					level[w.Vertex] = level_plus_one // Save level
				}
			}
		}
	}

	return parent, level
}

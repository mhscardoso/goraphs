package graphs

import (
	"github.com/mhscardoso/goraphs/queue"
	"github.com/mhscardoso/goraphs/stack"
)

// Executes the BFS
func BFS(g Graph, s int, know_signals *[]byte) ([]int, []uint, *queue.Queue) {
	vertices := g.N()

	if s <= 0 || s > vertices {
		return nil, nil, nil
	}

	// Signal to mark a vertex when discovered
	signal := know_signals
	if signal == nil {
		signals := make([]byte, vertices)
		signal = &signals
	}

	// parent[i] = j == vertex j is parent of i in BFS tree
	parent := make([]int, vertices)

	// Level of vertices in tree of BFS
	level := make([]uint, vertices)

	// Dealing with vertices between 0 -- N-1
	vertex := s - 1

	// Start a Queue
	Q := queue.New()

	// Storing the component in a Queue
	component := queue.New()

	// Mark given vertex and insert in queue
	(*signal)[vertex] = 1
	Q.Insert(vertex)

	component.Insert(vertex + 1)

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

				component.Insert(w.Vertex + 1) // Save vertex in the graph component
			}
		}
	}

	return parent, level, component
}

// Executes the DFS
func DFS(g Graph, s int, know_signals *[]byte) ([]int, []uint) {
	vertices := g.N()

	if s <= 0 || s > vertices {
		return nil, nil
	}

	// Working from 0 to N-1
	vertex := s - 1

	// To mark a vertex
	// Signal to mark a vertex when discovered
	signal := know_signals
	if signal == nil {
		signals := make([]byte, vertices)
		signal = &signals
	}

	// Parent os vertices in DFS
	parent := make([]int, vertices)

	// Level of vertices in DFS
	level := make([]uint, vertices)

	// Start a new stack and inserts the vertex in it
	P := stack.New()
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

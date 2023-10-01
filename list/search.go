package list

import (
	"github.com/mhscardoso/goraphs/queue"
	"github.com/mhscardoso/goraphs/stack"
)

func (L *List) BFS(s int) ([]byte, []int, []uint) {

	if s <= 0 || s > L.N {
		return nil, nil, nil
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

	return signal, parent, level
}

func (L *List) DFS(s int) ([]int, []uint) {

	if s <= 0 || s > L.N {
		return nil, nil
	}

	// Working from 0 to N-1
	vertex := s - 1

	// To mark a vertex
	signal := make([]byte, L.N)

	// Parent os vertices in DFS
	parent := make([]int, L.N)

	// Level of vertices in DFS
	level := make([]uint, L.N)

	// Start a new stack and inserts the vertex in it
	P := stack.New()
	P.Insert(vertex)

	// While P is not empty
	for e := P.Remove(); e != nil; e = P.Remove() {
		exploring := e.Vertex
		level_plus_one := level[exploring] + 1

		if signal[exploring] == 0 {
			signal[exploring] = 1

			for w := L.Vector[exploring].Next; w != nil; w = w.Next {
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

func (L *List) BFS_with_known_components(s int, signal []byte) ([]byte, *queue.Queue) {

	if s <= 0 || s > L.N {
		return nil, nil
	}

	// Dealing with vertices between 0 -- N-1
	vertex := s - 1

	// Start a Queue
	Q := queue.New()

	// Save vertices in current component
	Vertices_in_component := queue.New()

	// Mark given vertex and insert in queue
	signal[vertex] = 1
	Q.Insert(vertex)

	// For each vertex in queue; While queue not empty
	// It remeves a vertex in each iteration
	for e := Q.Remove(); e != nil; e = Q.Remove() {

		// Vertex being explored and its level in tree
		exploring := e.Vertex

		// For each neighbor of vertex
		for w := L.Vector[exploring].Next; w != nil; w = w.Next {
			if signal[w.Vertex] == 0 {
				signal[w.Vertex] = 1 // Mark vertex
				Q.Insert(w.Vertex)   // Insert in queue
				Vertices_in_component.Insert(w.Vertex)
			}
		}
	}

	return signal, Vertices_in_component
}

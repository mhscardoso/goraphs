package alist

// func (L *List) BFS_with_known_components(s int, signal *[]byte, component *queue.Queue) {
// 	if s <= 0 || s > L.N {
// 		return
// 	}

// 	// Dealing with vertices between 0 -- N-1
// 	vertex := s - 1

// 	// Start a Queue
// 	Q := queue.New()

// 	// Mark given vertex and insert in queue
// 	(*signal)[vertex] = 1
// 	Q.Insert(vertex)
// 	component.Insert(vertex)

// 	// For each vertex in queue; While queue not empty
// 	// It remeves a vertex in each iteration
// 	for e := Q.Remove(); e != nil; e = Q.Remove() {

// 		// Vertex being explored and its level in tree
// 		exploring := e.Vertex

// 		// For each neighbor of vertex
// 		for w := L.Vector[exploring].Next; w != nil; w = w.Next {
// 			if (*signal)[w.Vertex] == 0 {
// 				(*signal)[w.Vertex] = 1 // Mark vertex
// 				Q.Insert(w.Vertex)      // Insert in queue
// 				component.Insert(w.Vertex)
// 			}
// 		}
// 	}
// }

// func (L *List) Distance(s int, d int) int {

// 	if s <= 0 || s > L.N || d <= 0 || d > L.N {
// 		return -1
// 	}

// 	if s == d {
// 		return 0
// 	}

// 	// Signal to mark a vertex when discovered
// 	signal := make([]byte, L.N)

// 	// Level of vertices in tree of BFS
// 	level := make([]int, L.N)

// 	// Dealing with vertices between 0 -- N-1
// 	vertex := s - 1
// 	destiny := d - 1

// 	// Start a Queue
// 	Q := queue.New()

// 	// Mark given vertex and insert in queue
// 	signal[vertex] = 1
// 	Q.Insert(vertex)

// 	// For each vertex in queue; While queue not empty
// 	// It remeves a vertex in each iteration
// 	for e := Q.Remove(); e != nil; e = Q.Remove() {

// 		// Vertex being explored and its level in tree
// 		exploring := e.Vertex
// 		level_exploring_plus := level[exploring] + 1

// 		// For each neighbor of vertex
// 		for w := L.Vector[exploring].Next; w != nil; w = w.Next {
// 			if signal[w.Vertex] == 0 {
// 				signal[w.Vertex] = 1 // Mark vertex
// 				Q.Insert(w.Vertex)   // Insert in queue

// 				if w.Vertex == destiny {
// 					return level_exploring_plus
// 				}

// 				level[w.Vertex] = level_exploring_plus // Save its level
// 			}
// 		}
// 	}

// 	return -1
// }

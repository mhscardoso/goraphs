package matrix

import (
	"github.com/mhscardoso/goraphs/queue"
	"github.com/mhscardoso/goraphs/stack"
)

// s in [1, N]
// vertex in [0, N - 1]
func (A *AdjMatrix) BFS(s int) ([]byte, []int, []uint) {
	if s <= 0 || s > A.N {
		return nil, nil, nil
	}

	// Deal with vertices in matrix
	vertex := s - 1

	// Parent Vector
	parent := make([]int, A.N)

	// Level Vector
	level := make([]uint, A.N)

	// Where the vertices will be marked
	signal := make([]byte, A.N)
	signal[vertex] = 1

	// Queue to insert discovered vertices
	Q := queue.New()
	Q.Insert(vertex)

	// Important Data
	matrix_columns := A.N / BITS_IN_BYTE
	if A.N%BITS_IN_BYTE != 0 {
		matrix_columns++
	}

	for e := Q.TakeFirst(); e != nil; e = Q.TakeFirst() {
		exploring := e.Vertex
		level_exploring_plus := level[exploring] + 1
		Q.Remove()

		for j := 0; j < matrix_columns; j++ {
			supposed := A.P[exploring][j]

			// Iterating for bytes
			if supposed != byte(0) {

				// Iterating for the bits in byte
				for bit := 0; bit < 8; bit++ {

					// Checking bit by bit if is a neighbor
					if supposed&byte(1<<(7-bit)) != 0 {

						// If the neighbor not marked
						if signal[BITS_IN_BYTE*j+bit] != 1 {
							signal[BITS_IN_BYTE*j+bit] = 1 // Mark the neighbor
							Q.Insert(BITS_IN_BYTE*j + bit) // Insert thr neighbor in Q

							// here, remember, vertex = vertex in matrix + 1
							parent[BITS_IN_BYTE*j+bit] = exploring + 1       // exploring is parent from this vertex in tree
							level[BITS_IN_BYTE*j+bit] = level_exploring_plus // the level in tree is from exploring + 1
						}
					}
				}
			}
		}
	}
	return signal, parent, level
}

func (A *AdjMatrix) BFS_with_known_components(s int, signal *[]byte, component *queue.Queue) {
	if s <= 0 || s > A.N {
		return
	}

	// Deal with vertices in matrix
	vertex := s - 1
	(*signal)[vertex] = 1

	// Queue to insert discovered vertices
	Q := queue.New()
	Q.Insert(vertex)
	component.Insert(vertex)

	// Important Data
	matrix_columns := A.N / BITS_IN_BYTE
	if A.N%BITS_IN_BYTE != 0 {
		matrix_columns++
	}

	for e := Q.TakeFirst(); e != nil; e = Q.TakeFirst() {
		exploring := e.Vertex
		Q.Remove()

		for j := 0; j < matrix_columns; j++ {
			supposed := A.P[exploring][j]

			// Iterating for bytes
			if supposed != byte(0) {

				// Iterating for the bits in byte
				for bit := 0; bit < 8; bit++ {

					// Checking bit by bit if is a neighbor
					if supposed&byte(1<<(7-bit)) != 0 {

						// If the neighbor not marked
						if (*signal)[BITS_IN_BYTE*j+bit] != 1 {
							(*signal)[BITS_IN_BYTE*j+bit] = 1 // Mark the neighbor
							Q.Insert(BITS_IN_BYTE*j + bit)    // Insert thr neighbor in Q
							component.Insert(BITS_IN_BYTE*j + bit)
						}
					}
				}
			}
		}
	}
}

func (A *AdjMatrix) DFS(s int) ([]int, []uint) {
	if s <= 0 || s > A.N {
		return nil, nil
	}

	// Deal with vertices in matrix
	vertex := s - 1

	// Parent Vector
	parent := make([]int, A.N)

	// Level Vector
	level := make([]uint, A.N)

	// Where the vertices will be marked
	signal := make([]byte, A.N)

	// Queue to insert discovered vertices
	P := stack.New()
	P.Insert(vertex)

	// Important Data
	matrix_columns := A.N / BITS_IN_BYTE
	if A.N%BITS_IN_BYTE != 0 {
		matrix_columns++
	}

	for e := P.Remove(); e != nil; e = P.Remove() {
		exploring := e.Vertex
		level_plus_one := level[exploring] + 1

		if signal[exploring] != 1 {
			signal[exploring] = 1

			for j := 0; j < matrix_columns; j++ {
				supposed := A.P[exploring][j]

				// Iterating for bytes
				if supposed != byte(0) {

					// Iterating for the bits in byte
					for bit := 0; bit < 8; bit++ {

						// Checking bit by bit if is a neighbor
						if supposed&byte(1<<(7-bit)) != 0 {
							P.Insert(BITS_IN_BYTE*j + bit)

							if BITS_IN_BYTE*j+bit != vertex && parent[BITS_IN_BYTE*j+bit] == 0 {
								parent[BITS_IN_BYTE*j+bit] = exploring + 1
								level[BITS_IN_BYTE*j+bit] = level_plus_one
							}
						}
					}
				}
			}
		}
	}

	return parent, level
}

func (A *AdjMatrix) Distance(s int, d int) int {
	if s <= 0 || s > A.N || d <= 0 || d > A.N {
		return -1
	}

	// Deal with vertices in matrix
	vertex := s - 1
	destiny := d - 1

	// Level Vector
	level := make([]int, A.N)

	// Where the vertices will be marked
	signal := make([]byte, A.N)
	signal[vertex] = 1

	// Queue to insert discovered vertices
	Q := queue.New()
	Q.Insert(vertex)

	// Important Data
	matrix_columns := A.N / BITS_IN_BYTE
	if A.N%BITS_IN_BYTE != 0 {
		matrix_columns++
	}

	for e := Q.TakeFirst(); e != nil; e = Q.TakeFirst() {
		exploring := e.Vertex
		level_exploring_plus := level[exploring] + 1
		Q.Remove()

		for j := 0; j < matrix_columns; j++ {
			supposed := A.P[exploring][j]

			// Iterating for bytes
			if supposed != byte(0) {

				// Iterating for the bits in byte
				for bit := 0; bit < 8; bit++ {

					// Checking bit by bit if is a neighbor
					if supposed&byte(1<<(7-bit)) != 0 {

						// If the neighbor not marked
						if signal[BITS_IN_BYTE*j+bit] != 1 {
							signal[BITS_IN_BYTE*j+bit] = 1 // Mark the neighbor
							Q.Insert(BITS_IN_BYTE*j + bit) // Insert thr neighbor in Q

							// here, remember, vertex = vertex in matrix + 1
							if BITS_IN_BYTE*j+bit == destiny {
								return level_exploring_plus
							}

							level[BITS_IN_BYTE*j+bit] = level_exploring_plus // the level in tree is from exploring + 1
						}
					}
				}
			}
		}
	}
	return -1
}

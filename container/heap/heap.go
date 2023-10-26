package heap

import "math"

type E struct {
	Vertex   int
	Distance float64
}

// Min Heap implemented as a
// vector (or slice of int)
// in this case.
type Heap struct {
	Position []int
	Vector   []E
	Size     int
}

func New(capacity int) *Heap {
	heap := new(Heap)

	vector := make([]E, capacity)
	for i := range vector {
		vector[i].Distance = math.Inf(1)
		vector[i].Vertex = i
	}

	pos := make([]int, capacity)
	for i := range pos {
		pos[i] = i
	}

	heap.Vector = vector
	heap.Position = pos

	heap.Size = capacity

	return heap
}

// Utils
func LeftChildIndex(n int) int {
	return (2 * n) + 1
}
func RightChildIndex(n int) int {
	return (2 * n) + 2
}
func ParentIndex(n int) int {
	return (n - 1) / 2
}

func (h Heap) Parent(n int) float64 {
	return h.Vector[(n-1)/2].Distance
}

func (h Heap) LeftChild(n int) float64 {
	return h.Vector[(2*n)+1].Distance
}

func (h Heap) RightChild(n int) float64 {
	return h.Vector[(2*n)+2].Distance
}

func (h *Heap) Swap(i, j int) {
	h.Vector[i], h.Vector[j] = h.Vector[j], h.Vector[i]
	h.Position[h.Vector[i].Vertex], h.Position[h.Vector[j].Vertex] = h.Position[h.Vector[j].Vertex], h.Position[h.Vector[i].Vertex]
}

func (h *Heap) siftDown(current, destiny int) {
	leftChild := LeftChildIndex(current)
	for leftChild <= destiny {

		rightChild := RightChildIndex(current)

		if rightChild > destiny {
			rightChild = -1
		}

		// get the smaller child node to swap
		idxToSwap := leftChild
		if rightChild != -1 && h.Vector[rightChild].Distance < h.Vector[leftChild].Distance {
			idxToSwap = rightChild
		}

		// check if value of swap node is less than the value at currentIdx
		if h.Vector[idxToSwap].Distance < h.Vector[current].Distance {
			h.Swap(idxToSwap, current)
			current = idxToSwap
			leftChild = LeftChildIndex(current)

		} else {
			return
		}
	}
}

func (h *Heap) siftUp(i int) {
	current := i
	parent := ParentIndex(current)
	for current > 0 && h.Vector[current].Distance < h.Vector[parent].Distance {
		h.Swap(current, parent)
		current = parent
		parent = ParentIndex(current)
	}
}

func (h *Heap) Insert(Vertex int, Distance float64) {
	h.Vector = append(h.Vector, E{Vertex: Vertex, Distance: Distance})
	h.Position[Vertex] = h.Size
	h.Size++
	h.siftUp(h.Size - 1)
}

func (h *Heap) Remove() E {
	n := h.Size - 1
	// swap the first element and the last element in the array
	h.Swap(0, n)
	valueToRemove := h.Vector[n]
	h.Position[h.Vector[n].Vertex] = -1
	// pop the last element in the array
	h.Vector = h.Vector[:n]
	// call siftDown to update heap ordering
	h.Size--
	h.siftDown(0, n-1)

	return valueToRemove
}

func (h *Heap) Update(Vertex int, Distance float64) {
	heapPosition := h.Position[Vertex]

	old := h.Vector[heapPosition].Distance
	h.Vector[heapPosition].Distance = Distance

	if Distance < old {
		h.siftUp(heapPosition)
	} else if Distance > old {
		h.siftDown(0, heapPosition)
	}
}

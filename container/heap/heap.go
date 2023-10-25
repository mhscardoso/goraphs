package heap

type E struct {
	distance float64
	vertex   int
}

// Min Heap implemented as a
// vector (or slice of int)
// in this case.
type Heap struct {
	Vector []E
	Size   int
}

func New(capacity int) *Heap {
	heap := new(Heap)

	vector := make([]E, 0, capacity)

	heap.Vector = vector
	heap.Size = 0

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
	return h.Vector[(n-1)/2].distance
}

func (h Heap) LeftChild(n int) float64 {
	return h.Vector[(2*n)+1].distance
}

func (h Heap) RightChild(n int) float64 {
	return h.Vector[(2*n)+2].distance
}

func (h *Heap) Swap(i, j int) {
	h.Vector[i], h.Vector[j] = h.Vector[j], h.Vector[i]
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
		if rightChild != -1 && h.Vector[rightChild].distance < h.Vector[leftChild].distance {
			idxToSwap = rightChild
		}

		// check if value of swap node is less than the value at currentIdx
		if h.Vector[idxToSwap].distance < h.Vector[current].distance {
			h.Swap(idxToSwap, current)
			current = idxToSwap
			leftChild = LeftChildIndex(current)

		} else {
			return
		}
	}
}

func (h *Heap) siftUp() {
	current := h.Size - 1
	parent := ParentIndex(current)
	for current > 0 && h.Vector[current].distance < h.Vector[parent].distance {
		h.Swap(current, parent)
		current = parent
		parent = ParentIndex(current)
	}
}

func (h *Heap) Insert(vertex int, distance float64) {
	h.Vector = append(h.Vector, E{vertex: vertex, distance: distance})
	h.Size++
	h.siftUp()
}

func (h *Heap) Remove() E {
	n := h.Size - 1
	// swap the first element and the last element in the array
	h.Swap(0, n)
	valueToRemove := h.Vector[n]
	// pop the last element in the array
	h.Vector = h.Vector[:n]
	// call siftDown to update heap ordering
	h.Size--
	h.siftDown(0, n-1)

	return valueToRemove

}

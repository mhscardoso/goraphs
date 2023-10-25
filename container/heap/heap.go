package heap

type E struct {
	vertex   int
	distance float64
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

	vector := make([]E, 0, capacity)
	pos := make([]int, capacity)
	for i := range pos {
		pos[i] = -1
	}

	heap.Vector = vector
	heap.Position = pos

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
	h.Position[h.Vector[i].vertex], h.Position[h.Vector[j].vertex] = h.Position[h.Vector[j].vertex], h.Position[h.Vector[i].vertex]
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

func (h *Heap) siftUp(i int) {
	current := i
	parent := ParentIndex(current)
	for current > 0 && h.Vector[current].distance < h.Vector[parent].distance {
		h.Swap(current, parent)
		current = parent
		parent = ParentIndex(current)
	}
}

func (h *Heap) Insert(vertex int, distance float64) {
	h.Vector = append(h.Vector, E{vertex: vertex, distance: distance})
	h.Position[vertex] = h.Size
	h.Size++
	h.siftUp(h.Size - 1)
}

func (h *Heap) Remove() E {
	n := h.Size - 1
	// swap the first element and the last element in the array
	h.Swap(0, n)
	valueToRemove := h.Vector[n]
	h.Position[h.Vector[n].vertex] = -1
	// pop the last element in the array
	h.Vector = h.Vector[:n]
	// call siftDown to update heap ordering
	h.Size--
	h.siftDown(0, n-1)

	return valueToRemove
}

func (h *Heap) Update(vertex int, distance float64) {
	heapPosition := h.Position[vertex]

	// h.Vector[heapPosition].distance = math.Inf(1)
	// h.siftDown(heapPosition, h.Size-1)
	// if h.Vector[h.Size-1].distance != math.Inf(1) {
	// 	h.Swap(h.Size-1, h.Size-2)
	// }
	// h.Vector[h.Size-1].distance = distance
	// h.siftUp()

	old := h.Vector[heapPosition].distance
	h.Vector[heapPosition].distance = distance

	if distance < old {
		h.siftUp(heapPosition)
	} else if distance > old {
		h.siftDown(0, heapPosition)
	}
}

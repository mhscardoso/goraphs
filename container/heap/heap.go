package heap

// Min Heap implemented as a
// vector (or slice of int)
// in this case.
type Heap []int

func New(capacity int) *Heap {
	vector := make(Heap, capacity)
	return &vector
}

func (h Heap) Parent(n int) int {
	return h[(n-1)/2]
}

func (h Heap) LeftChild(n int) int {
	return h[(2*n)+1]
}

func (h Heap) RightChild(n int) int {
	return h[(2*n)+2]
}

package node

type Node[T interface{}] struct {
	Vertex T
	Next   *Node[T]
}

func NewNode[T interface{}](vertex T) *Node[T] {
	n := &Node[T]{Vertex: vertex, Next: nil}

	return n
}

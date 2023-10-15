package node

type Node struct {
	Vertex int
	Next   *Node
}

func NewNode(vertex int) *Node {
	n := &Node{Vertex: vertex, Next: nil}

	return n
}

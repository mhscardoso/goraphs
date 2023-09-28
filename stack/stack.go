package stack

import "fmt"

type Node struct {
	Vertex int
	Back   *Node
}

type Stack struct {
	Last  *Node
	First *Node
}

func New() *Stack {
	s := new(Stack)
	s.First = nil
	s.Last = s.First

	return s
}

func (s *Stack) Insert(vertex int) *Node {
	if s.Last == nil {
		first := new(Node)

		first.Vertex = vertex
		first.Back = nil

		s.First = first
		s.Last = first

		return first
	}

	top := s.First

	newNode := new(Node)
	newNode.Vertex = vertex
	newNode.Back = top

	s.First = newNode

	return newNode
}

func (s *Stack) Remove() *Node {
	removed := s.First
	if s.Last == s.First {
		s.First = nil
		s.Last = nil

		return removed
	}

	s.First = removed.Back
	return removed
}

func (s *Stack) See() {
	for e := s.First; e != nil; e = e.Back {
		fmt.Printf("%v\n", e.Vertex)
	}
}

func (s *Stack) Top() *Node {
	return s.First
}

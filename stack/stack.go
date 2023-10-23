package stack

import (
	"fmt"

	"github.com/mhscardoso/goraphs/lists"
	"github.com/mhscardoso/goraphs/node"
)

type Stack[T interface{}] struct {
	lists.Lists[T]
}

func New[T interface{}]() *Stack[T] {
	s := new(Stack[T])
	return s
}

func (s *Stack[T]) Insert(vertex T) *node.Node[T] {
	newNode := node.NewNode(vertex)
	if newNode == nil {
		fmt.Printf("Cannot allocate new node\n")
		return nil
	}

	if s.First == nil {
		s.First = newNode
		s.Last = newNode
	} else {
		newNode.Next = s.First
		s.First = newNode
	}

	s.Length++
	return newNode
}

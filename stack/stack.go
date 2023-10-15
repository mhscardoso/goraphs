package stack

import (
	"fmt"

	"github.com/mhscardoso/goraphs/lists"
	"github.com/mhscardoso/goraphs/node"
)

type Stack struct {
	lists.Lists
}

func New() *Stack {
	s := new(Stack)
	return s
}

func (s *Stack) Insert(vertex int) *node.Node {
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

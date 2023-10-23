package queue

import (
	"fmt"

	"github.com/mhscardoso/goraphs/lists"
	"github.com/mhscardoso/goraphs/node"
)

type Queue[T interface{}] struct {
	lists.Lists[T]
}

func New[T interface{}]() *Queue[T] {
	q := new(Queue[T])
	return q
}

func (q *Queue[T]) Insert(vertex T) *node.Node[T] {
	newNode := node.NewNode(vertex)
	if newNode == nil {
		fmt.Printf("Cannot allocate new node\n")
		return nil
	}

	if q.Last == nil {
		q.First = newNode
		q.Last = newNode
	} else {
		q.Last.Next = newNode
		q.Last = newNode
	}

	q.Length++
	return newNode
}

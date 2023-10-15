package queue

import (
	"fmt"

	"github.com/mhscardoso/goraphs/lists"
	"github.com/mhscardoso/goraphs/node"
)

type Queue struct {
	lists.Lists
}

func New() *Queue {
	q := new(Queue)
	return q
}

func (q *Queue) Insert(vertex int) *node.Node {
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

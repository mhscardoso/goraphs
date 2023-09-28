package queue

import "fmt"

type Node struct {
	Vertex int
	Back   *Node
}

type Queue struct {
	First *Node
	Last  *Node
}

func New() *Queue {
	q := new(Queue)
	q.First = nil
	q.Last = q.First

	return q
}

func (q *Queue) Insert(vertex int) *Node {
	if q.First == nil && q.Last == nil {
		first := new(Node)

		first.Vertex = vertex
		first.Back = nil

		q.First = first
		q.Last = first

		return first
	}

	last := q.Last

	newNode := new(Node)

	newNode.Vertex = vertex
	newNode.Back = nil

	q.Last = newNode
	last.Back = newNode

	return newNode
}

func (q *Queue) Remove() *Node {
	removed := q.First
	if q.Last == q.First {
		q.First = nil
		q.Last = nil

		return removed
	}

	q.First = removed.Back
	return removed
}

func (q *Queue) See() {
	for e := q.First; e != nil; e = e.Back {
		fmt.Printf("%v\n", e.Vertex)
	}
}

func (q *Queue) TakeFirst() *Node {
	return q.First
}

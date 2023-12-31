package lists

import (
	"fmt"

	"github.com/mhscardoso/goraphs/container/node"
)

type Lists[T interface{}] struct {
	First  *node.Node[T]
	Last   *node.Node[T]
	Length int
}

func (l Lists[T]) Top() *node.Node[T] {
	return l.First
}

func (l *Lists[T]) Remove() *node.Node[T] {
	removed := l.First

	if l.Last == l.First {
		l.First = nil
		l.Last = nil

		l.Length = 0
	} else {
		l.First = removed.Next
		l.Length--
	}

	return removed
}

func (l Lists[T]) See() {
	for e := l.First; e != nil; e = e.Next {
		fmt.Printf("%v\n", e.Vertex)
	}
}

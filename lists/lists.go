package lists

import (
	"fmt"

	"github.com/mhscardoso/goraphs/node"
)

type Lists struct {
	First  *node.Node
	Last   *node.Node
	Length int
}

func (l Lists) Top() *node.Node {
	return l.First
}

func (l *Lists) Remove() *node.Node {
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

func (l Lists) See() {
	for e := l.First; e != nil; e = e.Next {
		fmt.Printf("%v\n", e.Vertex)
	}
}

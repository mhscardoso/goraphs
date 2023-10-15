package sort

import (
	"github.com/mhscardoso/goraphs/lists"
	"github.com/mhscardoso/goraphs/stack"
)

func Sort(arr *[]int) {
	N := len(*arr)

	P := stack.New()

	P.Insert(0)
	P.Insert(N - 1)

	for e := P.Top(); e != nil; e = P.Top() {
		h := P.Remove().Vertex
		l := P.Remove().Vertex

		x := (*arr)[h]
		i := l - 1

		for j := l; j <= h-1; j++ {
			if (*arr)[j] <= x {
				i++
				Swap(&((*arr)[i]), &((*arr)[j]))
			}
		}
		Swap(&((*arr)[i+1]), &((*arr)[h]))

		if i > l {
			P.Insert(l)
			P.Insert(i)
		}

		if i+2 < h {
			P.Insert(i + 2)
			P.Insert(h)
		}
	}

}

func SortLists(arr []*lists.Lists, component int) {
	P := stack.New()

	P.Insert(0)
	P.Insert(component - 1)

	for e := P.Top(); e != nil; e = P.Top() {
		h := P.Remove().Vertex
		l := P.Remove().Vertex

		x := arr[h].Length
		i := l - 1

		for j := l; j <= h-1; j++ {
			if arr[j].Length <= x {
				i++
				SwapLists(arr[i], arr[j])
			}
		}
		SwapLists(arr[i+1], arr[h])

		if i > l {
			P.Insert(l)
			P.Insert(i)
		}

		if i+2 < h {
			P.Insert(i + 2)
			P.Insert(h)
		}
	}

}

func Swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func SwapLists(a *lists.Lists, b *lists.Lists) {
	t := *a
	*a = *b
	*b = t
}

package sort

import (
	"github.com/mhscardoso/goraphs/stack"
)

func Sort(arr *[]int) {
	N := len(*arr)

	P := stack.New[int]()

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

func Swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

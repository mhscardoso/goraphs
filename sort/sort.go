package sort

import (
	"github.com/mhscardoso/goraphs/queue"
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

func SortQueue(arr []*queue.Queue, component int) {
	P := stack.New()

	P.Insert(0)
	P.Insert(component - 1)

	for e := P.Top(); e != nil; e = P.Top() {
		h := P.Remove().Vertex
		l := P.Remove().Vertex

		x := arr[h].Size
		i := l - 1

		for j := l; j <= h-1; j++ {
			if arr[j].Size <= x {
				i++
				SwapQueues(arr[i], arr[j])
			}
		}
		SwapQueues(arr[i+1], arr[h])

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

func SwapQueues(a *queue.Queue, b *queue.Queue) {
	t := *a
	*a = *b
	*b = t
}

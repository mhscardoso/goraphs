package list

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Element struct {
	Vertex int
	Next   *Element
}

type List struct {
	Vector []Element
	N      int
	M      int
}

// Receives a Function
func CreateList(filename string, fn func(*Element, int) *Element) (*List, error) {
	readfile, err := os.Open(filename)

	if err != nil {
		readfile.Close()
		return nil, errors.New("Couldn`t read file")
	}

	// This happens when the function ends
	defer readfile.Close()

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanWords)

	// Number of Vertices
	// This is the first line in file
	fileScanner.Scan()
	v_s := fileScanner.Text()
	n, err := strconv.Atoi(v_s)

	if err != nil {
		readfile.Close()
		return nil, errors.New("Couldn`t convert first line")
	}

	// Creating Element Vector
	vec := make([]Element, n)

	// New list
	list := new(List)
	list.Vector = vec
	list.N = n

	// Numeber of edges
	m := 0

	for fileScanner.Scan() {
		// Each line, one edge
		m++

		v_p := fileScanner.Text()

		fileScanner.Scan()
		v_n := fileScanner.Text()

		vertex, err := strconv.Atoi(v_p)
		if err != nil {
			return nil, errors.New("One v_p could not be transformed into int")
		}

		neighbor, err := strconv.Atoi(v_n)
		if err != nil {
			return nil, errors.New("One v_n could not be transformed into int")
		}

		// Starting from 0
		vertex--
		neighbor--

		if vertex == neighbor {
			m--
			continue
		}

		// Add neighbor in order defined by fn - avoiding repeated
		newNeighbor := fn(&list.Vector[vertex], neighbor)

		if newNeighbor != nil {
			newNeighbor = fn(&list.Vector[neighbor], vertex)
		} else {
			m--
		}
	}

	list.M = m

	return list, nil
}

// Add a neighbor in order (to compare with matrix in BFS)
func AddSorted(e *Element, neighbor int) *Element {
	var it *Element
	for it = e; it.Next != nil; it = it.Next {
		if it.Next.Vertex > neighbor {
			break
		} else if it.Next.Vertex == neighbor {
			return nil
		}
	}

	newNeighbor := new(Element)
	newNeighbor.Vertex = neighbor

	next := it.Next
	it.Next = newNeighbor
	newNeighbor.Next = next

	return newNeighbor
}

// Add a neighbor in opposite order (to compare with matrix in DFS)
func AddNotSorted(e *Element, neighbor int) *Element {
	var it *Element
	for it = e; it.Next != nil; it = it.Next {
		if it.Next.Vertex < neighbor {
			break
		} else if it.Next.Vertex == neighbor {
			return nil
		}
	}

	newNeighbor := new(Element)
	newNeighbor.Vertex = neighbor

	next := it.Next
	it.Next = newNeighbor
	newNeighbor.Next = next

	return newNeighbor
}

func AddInOrder(e *Element, neighbor int) *Element {
	newNeighbor := new(Element)
	newNeighbor.Vertex = neighbor

	next := e.Next
	newNeighbor.Next = next
	e.Next = newNeighbor

	return newNeighbor
}

// Method to test in smalls graphs
func (l *List) See() {
	for i, v := range l.Vector {
		fmt.Printf("V: %v -- ", i+1)

		for e := v.Next; e != nil; e = e.Next {
			fmt.Printf("%v  ,  ", e.Vertex+1)
		}

		fmt.Printf("\n")
	}
}

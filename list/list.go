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

func CreateMatrix(filename string) (*List, error) {
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

		list.Vector[vertex].Add(neighbor)
		list.Vector[neighbor].Add(vertex)
	}

	list.M = m

	return list, nil
}

func (e *Element) Add(neighbor int) {
	var it *Element
	for it = e; it.Next != nil; it = it.Next {
		if it.Next.Vertex > neighbor {
			break
		}
	}

	newNeighbor := new(Element)
	newNeighbor.Vertex = neighbor

	next := it.Next
	it.Next = newNeighbor
	newNeighbor.Next = next
}

func (l *List) See() {
	for i, v := range l.Vector {
		fmt.Printf("V: %v -- ", i+1)

		for e := v.Next; e != nil; e = e.Next {
			fmt.Printf("%v  ,  ", e.Vertex+1)
		}

		fmt.Printf("\n")
	}
}

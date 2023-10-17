package graphs

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFile(g Graph, filename string) error {
	readfile, err := os.Open(filename)

	if err != nil {
		readfile.Close()
		return err
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
		return err
	}

	// Allocating Memory
	g.Allocate(n)

	// Numeber of edges
	m := 0

	for fileScanner.Scan() {
		v_p := fileScanner.Text()

		fileScanner.Scan()
		v_n := fileScanner.Text()

		vertex, err := strconv.Atoi(v_p)
		if err != nil {
			return err
		}

		neighbor, err := strconv.Atoi(v_n)
		if err != nil {
			return err
		}

		g.Relate(vertex, neighbor, &m)
	}

	g.UpdateEdges(m)

	return nil
}

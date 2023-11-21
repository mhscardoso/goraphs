package graphs

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/mhscardoso/goraphs/src/flist"
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
	fileScanner.Split(bufio.ScanLines)

	// Number of Vertices
	// This is the first line in file
	fileScanner.Scan()
	v_s := fileScanner.Text()
	n, err := strconv.Atoi(v_s)

	if err != nil {
		readfile.Close()
		return err
	}

	// Reading second line to check what
	// type of graph it is
	var fn func([]string) (int, int, float64)

	fileScanner.Scan()
	ln := fileScanner.Text()

	sl := strings.Split(ln, " ")
	if len(sl) == 3 {
		fn = Read3
	} else {
		fn = Read2
	}

	// Allocating Memory
	g.Allocate(n)

	// Numeber of edges
	m := 0

	for {
		vertex, neighbor, weigth := fn(sl)

		g.Relate(vertex, neighbor, weigth, &m)

		if !fileScanner.Scan() {
			break
		}

		ln = fileScanner.Text()
		sl = strings.Split(ln, " ")
	}

	g.UpdateEdges(m)

	return nil
}

// Read specifically graphs
// with integer flows
func ReadFileIntFlow(g *flist.FList, filename string) error {
	readfile, err := os.Open(filename)

	if err != nil {
		readfile.Close()
		return err
	}

	// This happens when the function ends
	defer readfile.Close()

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	// Number of Vertices
	// This is the first line in file
	fileScanner.Scan()
	v_s := fileScanner.Text()
	n, err := strconv.Atoi(v_s)

	if err != nil {
		readfile.Close()
		return err
	}

	// Reading second line to check what
	// type of graph it is
	var fn func([]string) (int, int, int)

	fileScanner.Scan()
	ln := fileScanner.Text()

	sl := strings.Split(ln, " ")
	if len(sl) == 3 {
		fn = Read3IntFlow
	}

	// Allocating Memory
	g.Allocate(n)

	// Numeber of edges
	m := 0

	for {
		vertex, neighbor, weigth := fn(sl)

		g.Relate(vertex, neighbor, weigth, &m)

		if !fileScanner.Scan() {
			break
		}

		ln = fileScanner.Text()
		sl = strings.Split(ln, " ")
	}

	g.UpdateEdges(m)

	return nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Read3(slc []string) (int, int, float64) {
	v, err := strconv.Atoi(slc[0])
	check(err)

	n, err := strconv.Atoi(slc[1])
	check(err)

	p, err := strconv.ParseFloat(slc[2], 32)
	check(err)

	return v, n, p
}

func Read3IntFlow(slc []string) (int, int, int) {
	v, err := strconv.Atoi(slc[0])
	check(err)

	n, err := strconv.Atoi(slc[1])
	check(err)

	p, err := strconv.Atoi(slc[2])
	check(err)

	return v, n, p
}

func Read2(slc []string) (int, int, float64) {
	v, err := strconv.Atoi(slc[0])
	check(err)

	n, err := strconv.Atoi(slc[1])
	check(err)

	return v, n, 0
}

func ReadFileNet(filename string) (map[int]string, map[string]int) {
	readfile, err := os.Open(filename)

	if err != nil {

		readfile.Close()
		return nil, nil
	}

	// This happens when the function ends
	defer readfile.Close()

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	mapIntString := make(map[int]string)
	mapStringInt := make(map[string]int)

	for fileScanner.Scan() {
		ln := fileScanner.Text()

		sl := strings.Split(ln, ",")

		vertex, err := strconv.Atoi(sl[0])

		if err != nil {
			return nil, nil
		}
		name := sl[1]

		mapIntString[vertex] = name
		mapStringInt[name] = vertex

	}
	return mapIntString, mapStringInt
}

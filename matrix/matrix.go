package matrix

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	BITS_IN_BYTE = 8
)

// N := number of vertices
// M := number of edges
// P := matrix per se
type AdjMatrix struct {
	N int
	M int
	P [][]byte
}

// Reading file: filename and creating
// the matrix of bits
func CreateMatrix(filename string) (*AdjMatrix, error) {
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

	// Will be stored N / 8 bytes, because
	// 1 byte has 8 bits. We will iterate over
	// bits. Thus, we need calculate the quotient
	// and remainder from division
	bytes := n / BITS_IN_BYTE
	rem := n % BITS_IN_BYTE
	if rem > 0 {
		bytes++
	}

	// Creating Matrix struct
	// M.P starts with nil (C NULL equivalent)
	Matrix := new(AdjMatrix)
	Matrix.N = n
	Matrix.P = make([][]byte, n)
	lines := make([]byte, n*bytes)

	for i := range Matrix.P {
		Matrix.P[i], lines = lines[:bytes], lines[bytes:]
	}

	// Number of Edges
	// Here, we realize the number of
	// Edges is the number of lines
	// in the file minus one
	m := 0

	// Read newxt lines of file
	for fileScanner.Scan() {
		// Each line, one edge!
		m++

		// Take two words per iteration (vertex and neighbor)
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

		// We want to start from zero in matrices and arrays
		vertex--
		neighbor--

		// Adding relationship in correct bit
		byte_house_one := neighbor / BITS_IN_BYTE
		byte_loc_one := (neighbor % BITS_IN_BYTE) + 1

		byte_house_two := vertex / BITS_IN_BYTE
		byte_loc_two := (vertex % BITS_IN_BYTE) + 1

		Matrix.P[vertex][byte_house_one] += byte(0b00000001 << (BITS_IN_BYTE - byte_loc_one))
		Matrix.P[neighbor][byte_house_two] += byte(0b00000001 << (BITS_IN_BYTE - byte_loc_two))
	}

	Matrix.M = m

	return Matrix, nil
}

// Method used to test small matrices
// It will print more numbers if we compare
// with the real matrix (you can use the test.txt graph to test!).
// This is caused by any remainder in the division of N by 8
// bytes can store 8 bits only, not less (nor more)
func (A *AdjMatrix) See() {
	fmt.Printf("    ")
	for i := 1; i <= A.N; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\n")

	for i := 0; i < A.N; i++ {
		fmt.Printf("%v - ", i+1)
		for j := 0; j < (A.N/BITS_IN_BYTE)+1; j++ {
			for k := BITS_IN_BYTE; k > 0; k-- {
				fmt.Printf("%v ", A.P[i][j]&byte(1<<(k-1))/byte(1<<(k-1)))
			}
		}
		fmt.Printf("\n")
	}
}

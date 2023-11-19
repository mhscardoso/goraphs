package main

import (
	"github.com/mhscardoso/goraphs/src/graphs"
)

func main() {
	A := graphs.FList(true)

	graphs.ReadFile(A, "testw.txt")

	A.See()
}

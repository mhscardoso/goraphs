package main

import (
	"fmt"

	"github.com/mhscardoso/goraphs/src/graphs"
)

func main() {
	A := graphs.FList(true)

	graphs.ReadFile(A, "testf.txt")

	b := graphs.Bottleneck(A, 1, 3)

	fmt.Printf("%v\n", b)
}

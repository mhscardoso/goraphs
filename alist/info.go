package alist

// import (
// 	"fmt"
// 	"os"

// 	"github.com/mhscardoso/goraphs/queue"
// 	"github.com/mhscardoso/goraphs/sort"
// )

// func (L *List) ConectedComponents() ([]*queue.Queue, int) {
// 	components := make([]*queue.Queue, L.N)
// 	component := 0
// 	signal := make([]byte, L.N)

// 	for i := 0; i < L.N; i++ {
// 		if signal[i] == 1 {
// 			continue
// 		}
// 		component++
// 		components[component-1] = queue.New()
// 		L.BFS_with_known_components(
// 			i+1,
// 			&signal,
// 			components[component-1])
// 	}

// 	sort.SortQueue(components, component)

// 	return components[:component], component
// }

// func (L *List) SaveData(filename string) {
// 	N_a, M_a, min_a, max_a, median_a, middle_a := L.GetInfo()

// 	f, err := os.Create(filename)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer f.Close()

// 	str := fmt.Sprintf("Vertices: %v\nEdges: %v\nMin Degree: %v\nMax Degree: %v\nMean Degrees: %v\nMedian Degrees: %v\n",
// 		N_a, M_a, min_a, max_a, median_a, middle_a)

// 	connected, g := L.ConectedComponents()

// 	str += fmt.Sprintf("Componentes Conexas: %v\n\n", g)

// 	for i := g - 1; i >= 0; i-- {
// 		str += fmt.Sprintf("Componente: %v - Tamanho: %v\n", i, connected[i].Size)

// 		for q := connected[i].Remove(); q != nil; q = connected[i].Remove() {
// 			str += fmt.Sprintf("%v ", q.Vertex)
// 		}

// 		str += fmt.Sprintf("\n\n")
// 	}

// 	fmt.Printf("Escrevendo o Arquivo\n")

// 	_, err2 := f.WriteString(str)
// 	if err2 != nil {
// 		panic(err2)
// 	}
// }
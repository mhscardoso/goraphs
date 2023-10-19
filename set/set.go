package set

import "fmt"

// Declaring the "void" type
type void struct{}

// The set type
type Set[K comparable] map[K]void

func New[K comparable]() Set[K] {
	s := Set[K]{}

	return s
}

func (s Set[K]) Add(e K) {
	s[e] = void{}
}

func (s Set[K]) Contains(e K) bool {
	_, ok := s[e]
	return ok
}

func (s Set[K]) Remove(e K) {
	if !s.Contains(e) {
		return
	}
	delete(s, e)
}

func (s Set[K]) See() {
	for i := range s {
		fmt.Printf("%v\n", i)
	}
	fmt.Printf("\n\n")
}

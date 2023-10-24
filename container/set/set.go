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

/*
 * Implementations for set specificaly to deal
 * weigth graphs.
 */

// The set type for graphs with weigths
type SetW[K comparable] map[K]float32

func NewW[K comparable]() SetW[K] {
	s := SetW[K]{}
	return s
}

func (s SetW[K]) Add(e K, w float32) {
	s[e] = w
}

func (s SetW[K]) Contains(e K) bool {
	_, ok := s[e]
	return ok
}

func (s SetW[K]) Remove(e K) {
	if !s.Contains(e) {
		return
	}
	delete(s, e)
}

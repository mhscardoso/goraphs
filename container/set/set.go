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
type SetW[K comparable, T float64 | int] map[K]T

func NewW[K comparable, T float64 | int]() SetW[K, T] {
	s := SetW[K, T]{}
	return s
}

func (s SetW[K, T]) Add(e K, w T) {
	s[e] = w
}

func (s SetW[K, T]) Contains(e K) bool {
	_, ok := s[e]
	return ok
}

func (s SetW[K, T]) Remove(e K) {
	if !s.Contains(e) {
		return
	}
	delete(s, e)
}

type waf [2]int

func (w *waf) SetWaf(weigth int, flow int) {
	w[0] = weigth
	w[1] = flow
}

func (w *waf) GetWeigth() int {
	return w[0]
}

func (w *waf) GetFlow() int {
	return w[1]
}

type SetF map[uint32]waf

func NewF() SetF {
	s := SetF{}
	return s
}

func (s SetF) Add(e uint32, w int, f int) {
	var data waf
	data.SetWaf(w, f)

	s[e] = data
}

func (s SetF) Contains(e uint32) bool {
	_, ok := s[e]
	return ok
}

func (s SetF) Remove(e uint32) {
	if !s.Contains(e) {
		return
	}
	delete(s, e)
}

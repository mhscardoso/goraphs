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
type SetW[K comparable] map[K]float64

func NewW[K comparable]() SetW[K] {
	s := SetW[K]{}
	return s
}

func (s SetW[K]) Add(e K, w float64) {
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

type waf [2]float64

func (w *waf) SetWaf(weigth float64, flow float64) {
	w[0] = weigth
	w[1] = flow
}

func (w *waf) GetWeigth() float64 {
	return w[0]
}

func (w *waf) GetFlow() float64 {
	return w[1]
}

type SetF map[uint32]waf

func NewF() SetF {
	s := SetF{}
	return s
}

func (s SetF) Add(e uint32, w float64, f float64) {
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

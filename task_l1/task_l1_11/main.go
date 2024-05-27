// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
	"strings"
)

// defining Set type as a map of empty structs
type Set[T comparable] struct {
	values map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		values: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(elem T) {
	s.values[elem] = struct{}{}
}

func (s1 *Set[T]) Intersect(s2 *Set[T]) *Set[T] {
	sRes := NewSet[T]()
	// finding the smallest set
	a, b := s1, s2
	if len(a.values) < len(b.values) {
		a, b = b, a
	}
	// iterating over the smallest set
	for key := range b.values {
		_, ok := a.values[key]
		if ok {
			sRes.Add(key)
		}
	}
	return sRes
}

// defining Set string representation
func (s *Set[T]) String() string {
	keys := make([]string, 0)
	for key, _ := range s.values {
		keys = append(keys, fmt.Sprint(key))
	}
	var str string = "Set [" + strings.Join(keys, " ") + "]"
	return str
}

func main() {
	s1 := NewSet[int]()
	for i := 0; i < 10; i++ {
		s1.Add(i)
	}

	s2 := NewSet[int]()
	for i := 5; i < 12; i++ {
		s2.Add(i)
	}

	s3 := s1.Intersect(s2)
	fmt.Printf("%v & %v = %v", s1, s2, s3)
}

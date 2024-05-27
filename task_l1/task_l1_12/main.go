// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

package main

import (
	"fmt"
	"strings"
)

// using Set type from task_11

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

// defining Set string representation
func (s *Set[T]) String() string {
	keys := make([]string, 0)
	for key := range s.values {
		keys = append(keys, fmt.Sprint(key))
	}
	var str string = "Set [" + strings.Join(keys, " ") + "]"
	return str
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSet[string]()
	for _, elem := range strings {
		set.Add(elem)
	}
	fmt.Print(set)
}

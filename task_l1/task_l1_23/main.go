package main

import (
	"errors"
	"fmt"
)

// deleting from slice with saving order
func deleteFromSlice1[T any](slice []T, i int) ([]T, error) {
	if i > len(slice)-1 {
		return nil, errors.New("Index out of range.\n")
	}
	slice = append(slice[:i], slice[i+1:]...)
	return slice, nil
}

// deleting from slice with rewriting an element with the last one
func deleteFromSlice2[T any](slice []T, i int) ([]T, error) {
	if i > len(slice)-1 {
		return nil, errors.New("Index out of range.\n")
	}
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1], nil
}

func main() {
	var s []int = make([]int, 10)
	fmt.Scan(s)

	var i int
	fmt.Scan(&i)

	s, _ = deleteFromSlice1(s, i)
	fmt.Println(s)
}

// Реализовать быструю сортировку массива (quicksort)
// встроенными методами языка.

package main

import "fmt"

// generic QSort implementation
func MyQSort[T any](s []T, less func(v1 T, v2 T) bool) {
	myQsort(s, 0, len(s)-1, less)
}

// rearranging elements all_left_elems < pivot < all_right elems
func partition[T any](s []T, left int, right int, less func(v1 T, v2 T) bool) int {
	pivotId := right
	wall := left
	for i := left; i < right; i++ {
		if less(s[i], s[pivotId]) {
			s[i], s[wall] = s[wall], s[i]
			wall++
		}
	}
	s[wall], s[right] = s[right], s[wall]
	return wall
}

func myQsort[T any](s []T, left int, right int, less func(v1 T, v2 T) bool) {
	if left < right {
		pivotId := partition(s, left, right, less)
		// sorting left part
		myQsort(s, left, pivotId-1, less)
		// sorting right part
		myQsort(s, pivotId+1, right, less)
	}
}

func main() {
	less := func(v1 int, v2 int) bool {
		return v1 < v2
	}

	s1 := []int{5, 7, 3, 7, 2, 8, 2, 5}
	fmt.Println(s1)
	MyQSort(s1, less)
	fmt.Println(s1)
	s2 := []int{1}
	fmt.Println(s2)
	MyQSort(s2, less)
	fmt.Println(s2)
	s3 := []int{2, 1}
	fmt.Println(s3)
	MyQSort(s3, less)
	fmt.Println(s3)

}

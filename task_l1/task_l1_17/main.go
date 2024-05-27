// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func BinarySearch[T any](
	s []T,
	v T,
	eq func(v1 T, v2 T) bool,
	less func(v1 T, v2 T) bool) int {
	l, r := 0, len(s)-1
	for l <= r {
		m := (l + r) / 2
		// checking mid element
		if eq(s[m], v) {
			return m
		}
		if less(s[m], v) {
			// moving left border
			l = m + 1
		} else {
			// moving right border
			r = m - 1
		}
	}
	// not found
	return -1
}

func main() {
	eq := func(v1 int, v2 int) bool {
		return v1 == v2
	}
	less := func(v1 int, v2 int) bool {
		return v1 < v2
	}

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := []int{1}
	s3 := []int{1, 2}

	fmt.Println(BinarySearch(s1, 2, eq, less))
	fmt.Println(BinarySearch(s1, 9, eq, less))
	fmt.Println(BinarySearch(s2, -1, eq, less))
	fmt.Println(BinarySearch(s2, 1, eq, less))
	fmt.Println(BinarySearch(s3, -1, eq, less))
	fmt.Println(BinarySearch(s3, 1, eq, less))
}

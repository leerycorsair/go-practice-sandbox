// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки
// должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	checks := make(map[rune]int)
	for _, r := range strings.ToLower(s) {
		_, ok := checks[r]
		if ok {
			return false
		}
		checks[r] = 1
	}
	return true
}

func main() {
	str1 := "abcd"
	fmt.Printf("%v\n", isUnique(str1))
	str2 := "abCdefAaf"
	fmt.Printf("%v\n", isUnique(str2))
	str3 := "aabcd"
	fmt.Printf("%v\n", isUnique(str3))
}

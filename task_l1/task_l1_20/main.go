// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func wordsReversedOrder(src string) string {
	words := strings.Split(src, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	srcStr := "lel kek cheburek rofl"
	newStr := wordsReversedOrder(srcStr)
	fmt.Println(newStr)
}

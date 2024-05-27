// Разработать программу, которая переворачивает подаваемую на
// ход строку (например: «главрыба — абырвалг»). Символы могут
// быть unicode.

package main

import "fmt"

func strReverse(src string) string {
	rvrStr := make([]rune, len(src))

	for i, r := range src {
		rvrStr[len(src)-1-i] = r
	}

	return string(rvrStr)
}

func main() {
	srcStr := "╭∩╮(︶︿︶)╭∩╮  🎀 KEK STRING 🎀  ╭∩╮(︶︿︶)╭∩╮"
	rvrStr := strReverse(srcStr)
	fmt.Print(rvrStr)
}

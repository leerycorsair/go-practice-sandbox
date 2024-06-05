package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackString(src string) (dst string, err error) {
	var previous rune
	escapeFlag := false

	for _, char := range src {
		if escapeFlag {
			dst += string(char)
			previous = char
			escapeFlag = false
		} else if char == '\\' {
			escapeFlag = true
		} else if unicode.IsDigit(char) {
			if previous == 0 {
				return "", errors.New("invalid str")
			}
			cnt, err := strconv.Atoi(string(char))
			if err != nil {
				return "", err
			}
			for i := 0; i < cnt-1; i++ {
				dst += string(previous)
			}
		} else {
			dst += string(char)
			previous = char
		}
	}

	if escapeFlag {
		return "", errors.New("invalid str")
	}

	return dst, nil
}

func main() {
	str := "a4bc2d5e"
	fmt.Println(UnpackString(str))
}

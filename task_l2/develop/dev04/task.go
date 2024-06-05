package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortedStr(src string) string {
	dst := strings.Split(src, "")
	sort.Strings(dst)
	return strings.Join(dst, "")
}

// Anagram is a function to find all sets of anagrams.
//
//	Input:[]string{"пятак", "пятка", "тяпка"}
//	Output:map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}
func Anagram(strs []string) map[string][]string {
	// forming a set of anagrams, where the key is a sorted string
	sortedKeyMap := make(map[string][]string, len(strs))
	addedWords := make(map[string]bool, len(strs))
	for _, str := range strs {
		lowerStr := strings.ToLower(str)
		// adding only unique strs
		if _, ok := addedWords[lowerStr]; !ok {
			key := sortedStr(lowerStr)
			addedWords[lowerStr] = true
			sortedKeyMap[key] = append(sortedKeyMap[key], lowerStr)
		}
	}

	// forming a set of anagrams, where the key is s 1st elem of each subset
	res := make(map[string][]string, len(strs))
	for _, strs := range sortedKeyMap {
		if len(strs) > 1 {
			sort.Strings(strs)
			res[strs[0]] = append(res[strs[0]], strs...)
		}
	}

	return res
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	ang := Anagram(words)

	fmt.Println(ang)
}

// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func main() {
// 	someFunc()
// }

// Строки в Go неизменяемые, при срезе строки создается новая строка, которая
// содержит ссылку на часть исходной строки. Это может привести к утечке памяти,
// если исходная строка очень большая, а нужна только её часть.

// Решением может быть создание новой строки и копирование в неё нужной части
// исходной строки. Это гарантирует, что память, занимаемая ненужными частями
// исходной строки, будет освобождена.

package main

func createHugeString(size int) string {
	return string(make([]byte, size))
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	tmp := make([]byte, 100)
	copy(tmp, v[:100])
	justString = string(tmp)
}

func main() {
	someFunc()
}

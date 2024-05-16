// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func swap1(num1, num2 *int) {
	// saving second value into first
	*num1 += *num2
	// setting first value as second
	*num2 = *num1 - *num2
	// setting second value as first
	*num1 -= *num2
}

func swap2(num1, num2 *int) {
	// saving second value into first
	*num1 ^= *num2
	// setting first value as second
	*num2 ^= *num1
	// setting second value as first
	*num1 ^= *num2
}

func main() {
	var num1, num2 int
	_, err := fmt.Scan(&num1, &num2)
	if err != nil {
		fmt.Printf("Incorrect input.\n")
		return
	}
	swap1(&num1, &num2)
	fmt.Println(num1, num2)
	swap2(&num1, &num2)
	fmt.Println(num1, num2)
}

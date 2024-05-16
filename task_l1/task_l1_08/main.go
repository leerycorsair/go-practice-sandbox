// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

func setBit(num *int64, bitId int64, bitValue int64) {
	if bitValue == 1 {
		// making i-bit as 1 and OR operatio
		*num |= (1 << bitId)
	} else {
		// making reverse mask and AND operation to reset i-bit
		*num &^= (1 << bitId)
	}
}

func main() {
	var number, bitId, bitValue int64
	_, err := fmt.Scan(&number, &bitId, &bitValue)
	if err != nil {
		fmt.Printf("Incorrect input.\n")
		return
	}
	setBit(&number, bitId, bitValue)
	fmt.Println(number)
}

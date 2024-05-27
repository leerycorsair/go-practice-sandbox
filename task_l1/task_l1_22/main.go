// Разработать программу, которая перемножает, делит,
// складывает, вычитает две числовых переменных a,b,
// значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	a.SetBit(a, 100, 1)
	b := big.NewInt(0)
	b.SetBit(b, 80, 1)

	buf := big.NewInt(0)
	fmt.Printf("sum = %s\n", buf.Add(a, b).Text(16))
	fmt.Printf("sub = %s\n", buf.Sub(a, b).Text(16))
	fmt.Printf("mul = %s\n", buf.Mul(a, b).Text(16))
	fmt.Printf("div = %s\n", buf.Div(a, b).Text(16))
}

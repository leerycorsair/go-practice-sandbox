// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

// using switch over type
func checkType1(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	case bool:
		fmt.Println("bool type")
	case chan int:
		fmt.Println("chan int type")
	case chan string:
		fmt.Println("chan string type")
	default:
		fmt.Println("unknown type")
	}
}

// using type assertions
func checkType2(v interface{}) {
	if _, ok := v.(int); ok {
		fmt.Println("int type")
		return
	}
	if _, ok := v.(string); ok {
		fmt.Println("string type")
		return
	}
	if _, ok := v.(bool); ok {
		fmt.Println("bool type")
		return
	}
	if _, ok := v.(chan int); ok {
		fmt.Println("chan int type")
		return
	}
	if _, ok := v.(chan string); ok {
		fmt.Println("chan string type")
		return
	}
	fmt.Println("unknown type")
}

func main() {
	var a int = 10
	var b string = "hello"
	var c bool = true
	var d chan int = make(chan int)
	var e chan string = make(chan string)
	var f float64 = 3.14

	checkType1(a)
	checkType1(b)
	checkType1(c)
	checkType1(d)
	checkType1(e)
	checkType1(f)

	checkType2(a)
	checkType2(b)
	checkType2(c)
	checkType2(d)
	checkType2(e)
	checkType2(f)
}

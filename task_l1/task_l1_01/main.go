// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) AgeInfo() string {
	return fmt.Sprintf("%s is %d years old.", h.Name, h.Age)
}

// Struct Embedding
type Action struct {
	Human
	Activity string
}

func (a Action) FullInfo() string {
	return a.AgeInfo() + fmt.Sprintf("%s will do %s.", a.Name, a.Activity)
}

func main() {
	a := &Action{
		Human: Human{
			Name: "Leonov Vlad",
			Age:  22,
		},
		Activity: "Testing",
	}
	fmt.Println(a.AgeInfo())
	fmt.Println(a.FullInfo())
}

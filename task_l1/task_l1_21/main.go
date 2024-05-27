// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// defining User that can say hello and goodbye
type User struct {
	name string
	a    UserActions
}

func (u *User) SayHello() {
	u.a.SayHello(u.name)
}

func (u *User) SayGoodBye() {
	u.a.SayGoodBye(u.name)
}

// interface with methods
type UserActions interface {
	SayHello(name string)
	SayGoodBye(name string)
}

// type that implements UserActions interface
type EngActions struct {
}

func (a *EngActions) SayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func (a *EngActions) SayGoodBye(name string) {
	fmt.Printf("Goodbye %s\n", name)
}

// type that doesn't implements UserActions interface
type RuActions struct {
}

func (a *RuActions) Privet(name []rune) {
	fmt.Printf("Привет %s\n", string(name))
}

func (a *RuActions) Poka(name []rune) {
	fmt.Printf("Пока %s\n", string(name))
}

// type that implements an adapter for RuActions to satisfy
// UserActions interface
type RuActionsAdapter struct {
	a *RuActions
}

func (aa *RuActionsAdapter) SayHello(name string) {
	aa.a.Privet([]rune(name))
}

func (aa *RuActionsAdapter) SayGoodBye(name string) {
	aa.a.Poka([]rune(name))
}

func main() {
	usr := User{name: "Vlad", a: &EngActions{}}
	usr.SayHello()
	usr.SayGoodBye()
	usr.a = &RuActionsAdapter{&RuActions{}}
	usr.SayHello()
	usr.SayGoodBye()
}

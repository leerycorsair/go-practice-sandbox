package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Применимость:
// 1) Когда объект сильно меняет свое поведение в зависимости от внутреннего состояния,
// состояний много и их действия меняются.
// 2) Когда код содержит большие условные операторы.
// 3) Когда используются условные операторы, но в некоторых блоках код дублируется.

// Плюсы:
// 1) Избавление от большых условных блоков.
// 2) Выделение кода для каждого состояния в отдельное место.
// 3) Упрощение кода контекста.

// Минусы:
// 1) Возможно усложнение кода, если состояний мало и они редко меняются.

type State interface {
	InsertCoin()
	PressButton()
	Dispense()
}

type CoffeeMachine struct {
	st State
}

func (m *CoffeeMachine) SetState(st State) {
	m.st = st
}

func (m *CoffeeMachine) InsertCoin() {
	m.st.InsertCoin()
}

func (m *CoffeeMachine) PressButton() {
	m.st.PressButton()
}

func (m *CoffeeMachine) Dispense() {
	m.st.Dispense()
}

type IdleState struct {
	m *CoffeeMachine
}

func (s *IdleState) InsertCoin() {
	fmt.Println("Coin inserted.")
	s.m.SetState(&BrewingState{s.m})
}

func (s *IdleState) PressButton() {
	fmt.Println("Insert coin first.")
}

func (s *IdleState) Dispense() {
	fmt.Println("Insert coin first.")
}

type BrewingState struct {
	m *CoffeeMachine
}

func (s *BrewingState) InsertCoin() {
	fmt.Println("Already brewing.")
}

func (s *BrewingState) PressButton() {
	fmt.Println("Brewing... Please wait.")
	s.m.SetState(&DispensingState{s.m})
}

func (s *BrewingState) Dispense() {
	fmt.Println("Brewing... Please wait.")
}

type DispensingState struct {
	m *CoffeeMachine
}

func (s *DispensingState) InsertCoin() {
	fmt.Println("Dispensing in progress.")
}

func (s *DispensingState) PressButton() {
	fmt.Println("Dispensing in progress.")
}

func (s *DispensingState) Dispense() {
	fmt.Println("Enjoy your coffee!")
	s.m.SetState(&IdleState{s.m})
}

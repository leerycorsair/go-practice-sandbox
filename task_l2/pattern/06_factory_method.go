package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Применимость:
// 1) Когда неизвестны типы и зависимости объектов, с которыми должен работать код.
// 2) Когда необходима возможность расширять части фреймворка/библиотеки.
// 3) Когда необходимо экономить системные ресурсы, используя повторно уже созданные объекты.

// Плюсы:
// 1) Отсутствие привязки класса к конкретным зависимым классам.
// 2) Явное выделение определения зависимостей в отдельное место.
// 3) Легкое добавление новых вариантов зависимостей.
// 4) Реализация принципа открытости/закрытости.

// Минусы:

type TravelAgent interface {
	SayHello()
}

type RuAgent struct {
}

func (a *RuAgent) SayHello() {
	fmt.Println("Привет")
}

type EngAgent struct {
}

func (a *EngAgent) SayHello() {
	fmt.Println("Hello")
}

type TravelAgency struct {
	agent TravelAgent
}

func NewTravelAgency(aType string) *TravelAgency {
	switch aType {
	case "eng":
		return &TravelAgency{&RuAgent{}}
	case "ru":
		return &TravelAgency{&EngAgent{}}
	default:
		return nil
	}
}

func (agency *TravelAgency) Welcome() {
	agency.agent.SayHello()
}

package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Применимость:
// 1) Когда нужно обрабатывать данные несколькими способами, но неизвестно
// какие будут данные и какие обработчики понадобятся.
// 2) Когда важно, чтобы обработчики выполнялись последовательно в строгом порядке.
// 3) Когда множество обработчиков является динамическим.

// Плюсы:
// 1) Уменьшение зависимости между клиентом и обработчиками.
// 2) Упрощение добавления новых обработчиков.
// 3) Логические разделение этапов вызова обработчиков.

// Минусы:
// 1) Данные могут быть остаться необработанными.

type Handler interface {
	SetNext(h Handler)
	Execute(data string)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(next Handler) {
	h.next = next
}

func (h *BaseHandler) Execute(data string) {
	if h.next != nil {
		h.next.Execute(data)
	}
}

type HandlerA struct {
	BaseHandler
}

func (h *HandlerA) Execute(data string) {
	if data == "a" {
		fmt.Println("HandlerA was executed")
	} else {
		fmt.Println("HandlerA was not executed, checking next...")
		h.BaseHandler.Execute(data)
	}
}

type HandlerB struct {
	BaseHandler
}

func (h *HandlerB) Execute(data string) {
	if data == "b" {
		fmt.Println("HandlerB was executed")
	} else {
		fmt.Println("HandlerB was not executed, checking next...")
		h.BaseHandler.Execute(data)
	}
}

type HandlerC struct {
	BaseHandler
}

func (h *HandlerC) Execute(data string) {
	if data == "c" {
		fmt.Println("HandlerC was executed")
	} else {
		fmt.Println("HandlerC was not executed, checking next...")
		h.BaseHandler.Execute(data)
	}
}

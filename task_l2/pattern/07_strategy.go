package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Применимость:
// 1) Когда нужно использовать разные вариации алгоритма внутри одного объекта.
// 2) Когда есть множество похожих классов, отличающихся только некоторым поведением.
// 3) Когда необходимо скрыть детали реализации алгоритма.

// Плюсы:
// 1) Динамическая смена алгоритмов.
// 2) Изоляция алгоритмов и их данных от остальных данных.
// 3) Уход от наследования к делегированию.

// Минусы:
// 1) Разрастание числа интерфейсов/типов.
// 2) Необходимо знать различия между стратегиями, чтобы выбрать подходящую.

type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCardStrategy struct {
	cardNumber string
}

func (s *CreditCardStrategy) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card ending in %s\n", amount, s.cardNumber[len(s.cardNumber)-4:])
}

type PayPalStrategy struct {
	email string
}

func (s *PayPalStrategy) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal account %s\n", amount, s.email)
}

type BitcoinStrategy struct {
	walletAddress string
}

func (s *BitcoinStrategy) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Bitcoin wallet %s\n", amount, s.walletAddress)
}

type PaymentProcessor struct {
	s PaymentStrategy
}

func (p *PaymentProcessor) SetStrategy(s PaymentStrategy) {
	p.s = s
}

func (p *PaymentProcessor) Pay(amount float64) {
	p.s.Pay(amount)
}

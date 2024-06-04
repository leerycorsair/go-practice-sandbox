package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Применимость:
// 1) Когда нужен простой/урезанный интерфейс для сложной подсистемы.
// 2) Когда нужно разделить большую подсистему на логические части.

// Плюсы:
// 1) Изоляция частей сложной подсистемы.
// 2) Упрощение взаимодействия с подсистемой.

// Минусы:
// 1) С усложнением ПО фасад может стать супер-объектом.

type OrderHandler struct {
	os *OrderService
	cs *CustomerService
	ss *SellerService
}

func (h *OrderHandler) CreateOrder() {
	fmt.Printf("Order Creating started\n")
	h.os.CheckOrder()
	h.cs.CheckCustomer()
	h.ss.MakeReservations()
	fmt.Printf("Order Creating finished\n")
}

type OrderService struct{}

func (s *OrderService) CheckOrder() (bool, error) {
	fmt.Printf("Checking Order Contents...\n")
	return true, nil
}

type CustomerService struct{}

func (s *CustomerService) CheckCustomer() (bool, error) {
	fmt.Printf("AntiBot system...\n")
	return true, nil
}

type SellerService struct{}

func (s *SellerService) MakeReservations() (bool, error) {
	fmt.Printf("Making Reservations of Items...\n")
	return true, nil
}

package service

import "module/internal/models"

type OrderService interface {
	CreateOrder(models.OrderT) (string, error)
	GetOrderByUid(string) (models.OrderT, error)
	GetOrders() ([]models.OrderT, error)
}

type Service struct {
	OrderService
}

package repository

import "module/internal/models"

type OrderRep interface {
	CreateOrder(o models.OrderT) (string, error)
	GetOrders() ([]models.OrderT, error)
	GetOrderByUid(uid string) (models.OrderT, error)
}

type Repository struct {
	OrderRep
}

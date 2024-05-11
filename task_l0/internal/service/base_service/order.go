package baseservice

import (
	"module/internal/models"
	"module/internal/repository"
)

type OrderService struct {
	rep *repository.Repository
}

func NewOrderService(rep *repository.Repository) *OrderService {
	return &OrderService{rep: rep}
}

func (s *OrderService) CreateOrder(o models.OrderT) (string, error) {
	return s.rep.CreateOrder(o)
}

func (s *OrderService) GetOrders() ([]models.OrderT, error) {
	return s.rep.GetOrders()
}

func (s *OrderService) GetOrderByUid(uid string) (models.OrderT, error) {
	return s.rep.GetOrderByUid(uid)
}

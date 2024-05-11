package baseservice

import (
	"module/internal/repository"
	"module/internal/service"
)

func NewBaseService(rep *repository.Repository) *service.Service {
	return &service.Service{
		OrderService: NewOrderService(rep),
	}
}

package dto

import "module/internal/models"

type DTOCreateOrder struct {
	OrderUid string `json:"order_uid"  db:"order_uid"`
}

type DTOGetOrders struct {
	Orders []models.OrderT `json:"orders"  db:"orders"`
}

type DTOGetOrder struct {
	Order models.OrderT `json:"order"  db:"order"`
}

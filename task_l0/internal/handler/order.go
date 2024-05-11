package handler

import (
	"module/internal/dto"
	"module/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary createOrder
// @Tags Order
// @Description Create a New Order
// @Accept json
// @Produce json
// @Param input body models.OrderT true "OrderInfo"
// @Success 200 {object} dto.DTOCreateOrder
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/orders/ [POST]
func (h *Handler) createOrder(c *gin.Context) {
	var order models.OrderT
	err := c.BindJSON(&order)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	orderUid, err := h.services.OrderService.CreateOrder(order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.cache.Add(orderUid, order)
	c.JSON(http.StatusOK, dto.DTOCreateOrder{OrderUid: orderUid})
}

// @Summary getOrder
// @Tags Order
// @Description Get All Orders
// @Produce json
// @Success 200 {object} dto.DTOGetOrders
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/orders/ [GET]
func (h *Handler) getOrders(c *gin.Context) {
	orders, err := h.services.GetOrders()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	for _, order := range orders {
		h.cache.Add(order.OrderUid, order)
	}
	c.JSON(http.StatusOK, dto.DTOGetOrders{Orders: orders})
}

// @Summary getOrderByUid
// @Tags Order
// @Description Get All Orders
// @Produce json
// @Param order_uid path string true "OrderUid"
// @Success 200 {object} dto.DTOGetOrder
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/orders/{order_uid} [GET]
func (h *Handler) getOrderByUid(c *gin.Context) {
	orderUid := c.Param("order_uid")
	order, found := h.cache.Get(orderUid)
	if !found {
		order, err := h.services.GetOrderByUid(orderUid)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, dto.DTOGetOrder{Order: order})
	}
	c.JSON(http.StatusOK, dto.DTOGetOrder{Order: order.(models.OrderT)})
}

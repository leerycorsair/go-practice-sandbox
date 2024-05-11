package handler

import (
	"io"
	"module/internal/cache"
	"module/internal/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"

	_ "module/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
	natsConn stan.Conn
	cache    *cache.Cache
}

func NewHandler(services *service.Service, natsConn stan.Conn, cache *cache.Cache) *Handler {
	return &Handler{services: services, natsConn: natsConn, cache: cache}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		orders := api.Group("/orders")
		{
			orders.POST("/", h.createOrder)
			orders.GET("/", h.getOrders)
			orders.GET("/:order_uid", h.getOrderByUid)
		}
		api.POST("/publish", h.publishMsg)
	}
	return router
}

// @Summary publishMsg
// @Tags Msg
// @Description Publish message
// @Accept json
// @Produce json
// @Param message body string true "Message content to publish"
// @Success 200 {integer} integer
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/publish [POST]
func (h *Handler) publishMsg(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.natsConn.Publish(viper.GetString("nats.client_id"), jsonData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "message was published")

}

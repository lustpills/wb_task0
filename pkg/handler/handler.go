package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	orders := router.Group("/orders")
	{
		orders.POST("/", h.createOrder)
		orders.GET("/:id", h.getOrderById)
	}
	return router
}

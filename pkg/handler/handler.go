package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lustpills/wb_task0/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/css", "./templates/css")
	router.LoadHTMLGlob("templates/*.html")

	orders := router.Group("/orders")
	{
		orders.GET("/", h.createOrder)
		orders.GET("/:id", h.getOrderById)
	}
	return router
}

package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	caching "github.com/lustpills/wb_task0/pkg/cashing"
)

func (h *Handler) createOrder(c *gin.Context) {
	//var input orders.Order

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(c.Writer, nil)
}

func (h *Handler) getOrderById(c *gin.Context) {

	input := c.Param("id")
	fmt.Println(input)

	// tmpl := template.Must(template.ParseFiles("index.html"))
	// tmpl.Execute(c.Writer, nil)

	//order, err := h.services.GetOrder(input)
	order, err := caching.MyCache.Get(input)

	if !err {
		newErrorResponse(c, http.StatusInternalServerError, "order not found")
		return
	}

	//order_json, _ := json.Marshal(order)

	c.JSON(http.StatusOK, order)
}

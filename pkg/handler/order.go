package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	caching "github.com/lustpills/wb_task0/pkg/cashing"
)

func (h *Handler) createOrder(c *gin.Context) {

	// main site page handler
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(c.Writer, nil)
}

func (h *Handler) getOrderById(c *gin.Context) {

	input := c.Param("id")
	fmt.Println(input)

	// getting order detales from cache by an order_uid key
	order, err := caching.MyCache.Get(input)

	if !err {
		newErrorResponse(c, http.StatusInternalServerError, "order not found")
		//c.JSON(http.StatusNoContent, "order not found")
	} else {
		// sending order detales
		c.JSON(http.StatusOK, order)
	}

}

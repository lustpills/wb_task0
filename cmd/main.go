package main

import (
	"log"

	orders "github.com/lustpills/wb_task0"
	"github.com/lustpills/wb_task0/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(orders.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while runnint http server: ", err)
	}
}

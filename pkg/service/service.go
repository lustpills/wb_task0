package service

import (
	orders "github.com/lustpills/wb_task0"
	"github.com/lustpills/wb_task0/pkg/repository"
)

type Orders interface {
	CreateOrder(order orders.Order) (string, error)
	RestoreCache() (map[string]interface{}, error)
	GetOrder(string) (orders.Order, error)
	//CacheOrder(order)
}

type Service struct {
	Orders
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Orders: NewOrderService(repos.Orders),
	}
}

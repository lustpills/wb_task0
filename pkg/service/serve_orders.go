package service

import (
	orders "github.com/lustpills/wb_task0"
	"github.com/lustpills/wb_task0/pkg/repository"
)

type OrderService struct {
	repo repository.Orders
}

func NewOrderService(repo repository.Orders) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order orders.Order) (string, error) {
	return s.repo.CreateOrder(order)
}

func (s *OrderService) RestoreCache() (map[string]interface{}, error) {
	return s.repo.RestoreCache()
}

func (s *OrderService) GetOrder(order_id string) (orders.Order, error) {
	return s.repo.GetOrder(order_id)
}

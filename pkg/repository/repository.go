package repository

import (
	"github.com/jmoiron/sqlx"
	orders "github.com/lustpills/wb_task0"
)

type Orders interface {
	CreateOrder(order orders.Order) (string, error)
	RestoreCache() (map[string]interface{}, error)
	GetOrder(order_id string) (orders.Order, error)
}

type Repository struct {
	Orders
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Orders: NewOrderPostgres(db),
	}
}

package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	orders "github.com/lustpills/wb_task0"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

// CreateOrder stores order in a db
func (r *OrderPostgres) CreateOrder(order orders.Order) (string, error) {
	var order_id string

	query := fmt.Sprintf("INSERT INTO %s (order_uid, track_number"+
		", entry, delivery, payment, items, locale, internal_signature, customer_id"+
		", delivery_service, shardkey, sm_id, date_created, oof_shard) "+
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING order_uid", ordersTable)

	row := r.db.QueryRow(query, order.Order_uid, order.Track_number, order.Entry, order.Delivery,
		order.Payment, order.Items, order.Locale, order.Internal_signature, order.Customer_id,
		order.Delivery_service, order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard)

	if err := row.Scan(&order_id); err != nil {
		return "", err
	}
	return order_id, nil
}

// RestoreCache transforms db into a map and returns it
func (r *OrderPostgres) RestoreCache() (map[string]interface{}, error) {
	RestoredSlice := []orders.Order{}
	RestoredCache := make(map[string]interface{})

	query := fmt.Sprintf("SELECT * FROM %s", ordersTable)
	err := r.db.Select(&RestoredSlice, query)

	if err != nil {
		log.Fatal("Error occured while restoring cache: ", err)
	}

	for i := 0; i < len(RestoredSlice); i++ {
		uid := string(RestoredSlice[i].Order_uid)
		fmt.Println("Cache: ")
		fmt.Println(uid)
		RestoredCache[uid] = RestoredSlice[i]
	}

	return RestoredCache, nil
}

// GetOrder get an order from a db by its order_uid primary key
func (r *OrderPostgres) GetOrder(order_id string) (orders.Order, error) {

	var order orders.Order

	query := fmt.Sprintf("SELECT * FROM %s WHERE order_uid=$1", ordersTable)

	err := r.db.Get(&order, query, order_id)
	return order, err
}

package publish

import (
	"encoding/json"
	"math/rand"
	"time"

	orders "github.com/lustpills/wb_task0"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

// generate returns a marshaled json model of an order with random detales

func generate() []byte {
	rand.Seed(time.Now().UnixNano())

	var random_order orders.Order

	json.Unmarshal([]byte(Msg_j), &random_order)

	b := make([]rune, 13)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	random_order.Order_uid = string(b)

	random_order_json, _ := json.Marshal(random_order)
	return random_order_json
}

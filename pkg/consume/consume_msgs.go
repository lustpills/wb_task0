package consume

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	orders "github.com/lustpills/wb_task0"
	caching "github.com/lustpills/wb_task0/pkg/cashing"
	"github.com/nats-io/nats.go/jetstream"
)

func (c *Consumer) Consuming(ctx context.Context) {
	var new_order orders.Order

	// creating a consumer to a stream (practically same as subscriber)
	NewConsumer, err := c.stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "CONS",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatal("Error occured while creating a consumer: ", err)
	}

	fmt.Println("consumer created")

	messageCounter := 0

	// reading orders from the stream and acknoledging them
	for {
		msgs, _ := NewConsumer.Fetch(10)
		for msg := range msgs.Messages() {
			json.Unmarshal(msg.Data(), &new_order)

			// first putting order in db
			order_id, err := c.services.Orders.CreateOrder(new_order)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println("Sucsessfully stored an order: ", order_id)
				caching.MyCache.Set(order_id, new_order)
			}

			// second caching an order
			// /caching.MyCache.Set(order_id, new_order)
			messageCounter++
			msg.Ack()
		}
	}

}

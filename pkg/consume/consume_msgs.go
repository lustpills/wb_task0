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
	NewConsumer, err := c.stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "CONS",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatal("Error occured while creating a consumer: ", err)
	}

	fmt.Println("consumer created")

	messageCounter := 0
	//msgs, _ := NewConsumer.FetchNoWait(10)
	for {
		msgs, _ := NewConsumer.Fetch(10)
		for msg := range msgs.Messages() {

			//fmt.Printf("Received a JetStream message via fetch: %s\n", string(msg.Data()))
			json.Unmarshal(msg.Data(), &new_order)
			order_id, err := c.services.Orders.CreateOrder(new_order)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println("Sucsessfully stored an order: ", order_id)
			}
			caching.MyCache.Set(order_id, new_order)
			messageCounter++
			msg.Ack()
		}
		//_ = c.stream.DeleteMsg(ctx, 0)
	}

	// cons, err := NewConsumer.Consume(func(msg jetstream.Msg) {
	// 	msg.Ack()
	// 	data := msg.Data()
	// 	fmt.Printf("Received a JetStream message via callback: %s\n", string(data))
	// 	json.Unmarshal(data, &new_order)

	// 	order_id, err := c.services.Orders.CreateOrder(new_order)

	// 	if err != nil {
	// 		log.Println("Error occured while creating an order: ", order_id, err)
	// 	}

	// 	//messageCounter++

	// })
	// if err != nil {
	// 	log.Fatal("Error occured while consuming messages: ", err)
	// }
	// defer cons.Stop()
}

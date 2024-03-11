package publish

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lustpills/wb_task0/pkg/service"
	"github.com/nats-io/nats.go/jetstream"
)

type Publisher struct {
	services *service.Service
}

func NewPublisher(services *service.Service) *Publisher {
	return &Publisher{services: services}
}

func (p *Publisher) Publishing(ctx context.Context, s jetstream.Stream, js jetstream.JetStream) {

	msg_counter := 0

	for {
		time.Sleep(2 * time.Hour)
		_, err := js.Publish(ctx, "ORDERS.new", generate())
		if err != nil {
			log.Fatal("error occured while trying to publish an order: ", err)
		}
		fmt.Printf("Published hello message %d\n", msg_counter)
		msg_counter++

	}
	//msg, _ := s.GetLastMsgForSubject(ctx, "ORDERS.new")
	//var new_order orders.Order

	//json.Unmarshal(msg.Data, &new_order)

	//fmt.Println(new_order.Order_uid)

	//order_id, err := p.services.Orders.CreateOrder(new_order)

	// if err != nil {
	// 	log.Fatal("error occured while creating an order: ", err)
	// }

	//fmt.Println(order_id)

}

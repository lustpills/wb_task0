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

// publish orders in a cycle, shold be started with a goroutine
func (p *Publisher) Publishing(ctx context.Context, s jetstream.Stream, js jetstream.JetStream) {

	msg_counter := 0

	for {
		time.Sleep(10 * time.Second)
		_, err := js.Publish(ctx, "ORDERS.new", generate())
		if err != nil {
			log.Fatal("error occured while trying to publish an order: ", err)
		}
		fmt.Printf("Published hello message %d\n", msg_counter)
		msg_counter++

	}

}

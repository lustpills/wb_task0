package consume

import (
	"github.com/lustpills/wb_task0/pkg/service"
	"github.com/nats-io/nats.go/jetstream"
)

type Consumer struct {
	services *service.Service
	stream   jetstream.Stream
}

func NewConsumer(services *service.Service, s jetstream.Stream) *Consumer {
	return &Consumer{services: services, stream: s}
}

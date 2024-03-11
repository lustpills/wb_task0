package consuming

import (
	"log"

	"github.com/nats-io/nats.go"
)

type MessageHandler interface {
	HandleMessage(msg *nats.Msg)
}

type NatsConsumerService struct {
	jsContext      nats.JetStreamContext
	subject        string
	messageHandler MessageHandler
}

func NewNatsConsumerService(jsContext nats.JetStreamContext, subject string, messageHandler MessageHandler) *NatsConsumerService {
	return &NatsConsumerService{
		jsContext:      jsContext,
		subject:        subject,
		messageHandler: messageHandler,
	}
}

func (service *NatsConsumerService) StartListening() {
	_, err := service.jsContext.Subscribe(service.subject, func(msg *nats.Msg) {
		service.messageHandler.HandleMessage(msg)
	})

	if err != nil {
		log.Fatalf("Failed to subscribe to subject %s: %v", service.subject, err)
	}

	log.Printf("Subscribed to %s", service.subject)
}

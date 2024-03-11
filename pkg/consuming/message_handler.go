package consuming

import (
	"encoding/json"
	"log"

	orders "github.com/lustpills/wb_task0"
	caching "github.com/lustpills/wb_task0/pkg/cashing"
	"github.com/lustpills/wb_task0/pkg/service"
	"github.com/nats-io/nats.go"
)

type CacheMessageHandler struct {
	cache    *caching.Cache
	services *service.Service
}

func NewCacheMessageHandler(cache *caching.Cache, services *service.Service) *CacheMessageHandler {
	return &CacheMessageHandler{cache: cache, services: services}
}

func (handler *CacheMessageHandler) HandleMessage(msg *nats.Msg) {
	// Десериализация сообщения из msg.Data
	// Предположим, что сообщения передаются в формате JSON
	var data orders.Order
	err := json.Unmarshal(msg.Data, &data)
	if err != nil {
		log.Printf("Ошибка десериализации сообщения: %v", err)
		// Негативное подтверждение сообщения, если необходимо
		// msg.Nak()
		return
	}

	// Обработка данных из сообщения, например, сохранение в кэш или БД
	// Здесь мы вызываем метод 'Set' нашего кэша, предполагая, что он принимает ключ и значение.
	// В качестве ключа используем, например, уникальный идентификатор из данных сообщения.
	order_id, err := handler.services.Orders.CreateOrder(data)
	if err != nil {
		log.Println(err, order_id)
		return
	}
	handler.cache.Set(data.Order_uid, data)

	// Подтверждение успешной обработки сообщения.
	// Это сообщит NATS Jetstream, что сообщение было обработано и его можно удалить из стрима.
	err = msg.Ack()
	if err != nil {
		log.Printf("Ошибка подтверждения сообщения: %v", err)
		return
	}
}

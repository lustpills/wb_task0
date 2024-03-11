package main

import (
	"context"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"

	orders "github.com/lustpills/wb_task0"
	caching "github.com/lustpills/wb_task0/pkg/cashing"
	"github.com/lustpills/wb_task0/pkg/consume"
	"github.com/lustpills/wb_task0/pkg/handler"
	"github.com/lustpills/wb_task0/pkg/publish"
	"github.com/lustpills/wb_task0/pkg/repository"
	"github.com/lustpills/wb_task0/pkg/service"
	"github.com/spf13/viper"
)

const (
	msg_j = `{
		"order_uid": "b563feb7b2b84b6test",
		"track_number": "WBILMTESTTRACK",
		"entry": "WBIL",
		"delivery": {
		  "name": "Test Testov",
		  "phone": "+9720000000",
		  "zip": "2639809",
		  "city": "Kiryat Mozkin",
		  "address": "Ploshad Mira 15",
		  "region": "Kraiot",
		  "email": "test@gmail.com"
		},
		"payment": {
		  "transaction": "b563feb7b2b84b6test",
		  "request_id": "",
		  "currency": "USD",
		  "provider": "wbpay",
		  "amount": 1817,
		  "payment_dt": 1637907727,
		  "bank": "alpha",
		  "delivery_cost": 1500,
		  "goods_total": 317,
		  "custom_fee": 0
		},
		"items": [
		  {
			"chrt_id": 9934930,
			"track_number": "WBILMTESTTRACK",
			"price": 453,
			"rid": "ab4219087a764ae0btest",
			"name": "Mascaras",
			"sale": 30,
			"size": "0",
			"total_price": 317,
			"nm_id": 2389212,
			"brand": "Vivienne Sabo",
			"status": 202
		  }
		],
		"locale": "en",
		"internal_signature": "",
		"customer_id": "test",
		"delivery_service": "meest",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "2021-11-26T06:22:19Z",
		"oof_shard": "1"
	  }`
)

func main() {

	//var new_order orders.Order
	//new_order := new(orders.Order)
	//json.Unmarshal([]byte(msg_j), &new_order)

	//fmt.Println(new_order.Date_created)

	if err := initConfig(); err != nil {
		log.Fatal("error initializing configs: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("failed to nats ", err)
	}
	// Create a JetStream management interface
	js, _ := jetstream.New(nc)

	// Create a stream
	s, _ := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})

	// Publish some messages
	// for i := 0; i < 10; i++ {
	// 	js.Publish(ctx, "ORDERS.new", []byte(msg_j))
	// 	fmt.Printf("Order created %d\n", i)
	// }

	// msg, _ := s.GetLastMsgForSubject(ctx, "ORDERS.new")

	// json.Unmarshal(msg.Data, &new_order)

	// fmt.Println(new_order.Order_uid)

	// Create durable consumer
	// c, _ := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
	// 	Durable:   "CONS",
	// 	AckPolicy: jetstream.AckExplicitPolicy,
	// })

	// Get 10 messages from the consumer
	// messageCounter := 0
	// msgs, _ := c.Fetch(10)
	// for msg := range msgs.Messages() {
	// 	msg.Ack()
	// 	fmt.Printf("Received a JetStream message via fetch: %s\n", string(msg.Data()))
	// 	messageCounter++
	// }

	// for i := 1; i <= 1000; i++ {
	// 	sc.Publish("besellungen", []byte("Bestellung "+strconv.Itoa(i)))
	// }

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "hell",
		DBname:   "orders",
		SSlmode:  "disable",
	})
	if err != nil {
		log.Fatal("failed to init db: ", err)
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	publishers := publish.NewPublisher(service)
	consumers := consume.NewConsumer(service, s)
	cache := caching.NewCash(service)

	//consuming.MessageHandler
	//OrderConsumer := consuming.NewNatsConsumerService(js, "ORDERS", cache)

	//go publishers.Publishing(ctx, s, js)
	go publishers.Publishing(ctx, s, js)
	// go publishers.Publishing(ctx, s, js)
	// go publishers.Publishing(ctx, s, js)
	// go publishers.Publishing(ctx, s, js)
	// go publishers.Publishing(ctx, s, js)

	go consumers.Consuming(ctx)

	cache.Restore()

	//msg, _ := s.GetLastMsgForSubject(ctx, "ORDERS.new")
	//cache.Set()

	srv := new(orders.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while runnint http server: ", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

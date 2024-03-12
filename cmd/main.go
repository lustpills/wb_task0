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

func main() {

	if err := initConfig(); err != nil {
		log.Fatal("error initializing configs: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// initiating a nats-server connection
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

	// initiating a db to store orders
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

	// starting a publishing goroutine
	go publishers.Publishing(ctx, s, js)

	// starting stream listening goroutine
	go consumers.Consuming(ctx)

	cache.Restore()

	// starting a server
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

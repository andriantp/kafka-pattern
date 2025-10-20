package main

import (
	"context"
	"kafka-pattern/adapter/kafka"
	"kafka-pattern/logger"
)

func main() {
	ctx := context.Background()
	conf := kafka.Config{
		Broker:  "kafka:9092",
		Topic:   "dbserver1.public.products",
		GroupID: "group-golang1",
	}

	consumer, err := kafka.NewConsumer(conf)
	if err != nil {
		logger.Level("fatal", "main:NewConsumer", err)
	}
	consumer.Consume(ctx)
}

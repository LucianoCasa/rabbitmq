package main

import (
	rabbitmq "rabbitmq/pkg"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	mq, err := rabbitmq.NewMQ("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer mq.Close()

	mq.Publish("teste", "Hello World!")

	msgs := make(chan amqp.Delivery)

	go mq.Consume("teste", msgs)

	for msg := range msgs {
		println("Received message:", string(msg.Body))
	}
}

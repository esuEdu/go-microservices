package main

import (
	"fmt"
	"log"
	"os"
	"payment-service/internal/processors"
	"payment-service/internal/rabbit"
)

func main() {

	rabbitURL := fmt.Sprintf(
		"amqp://%s:%s@%s:5672/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASS"),
		os.Getenv("RABBITMQ_HOST"),
	)

	rc, err := rabbit.NewRabbitConnection(rabbitURL)
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	publisher := rabbit.NewPublisher(rc.Ch)

	processor := processors.NewPaymentProcessor(publisher)

	consumer := rabbit.NewConsumer(processor)

	consumer.Consume(rc, "order.created")

	select {}
}

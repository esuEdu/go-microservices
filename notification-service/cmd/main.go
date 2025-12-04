package main

import (
	"fmt"
	"log"
	"notification-service/internal/notifier"
	"notification-service/internal/rabbit"
	"os"
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
		log.Fatal("RabbitMQ connection failed:", err)
	}
	defer rc.Close()

	consumer := rabbit.NewPaymentConsumer(notifier.SendEmail)

	log.Println("📨 Notification Service waiting for payment.completed...")
	err = consumer.Consume(rc, "payment.completed")
	if err != nil {
		log.Fatal(err)
	}

	select {}
}

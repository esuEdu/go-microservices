package main

import (
	"fmt"
	"log"
	"net/http"
	"order-service/internal/api"
	"order-service/internal/rabbit"
	"os"
)

func main() {

	rabbitURL := fmt.Sprintf(
		"amqp://%s:%s@%s:5672/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASS"),
		os.Getenv("RABBITMQ_HOST"),
	)

	err := rabbit.ConnectRabbitMQ(rabbitURL)
	if err != nil {
		log.Fatal("RabbitMQ connection error:", err)
	}
	defer rabbit.CloseRabbit()

	router := api.SetupRouter()

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", router)
}

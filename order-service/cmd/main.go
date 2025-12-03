package main

import (
	"fmt"
	"log"
	"net/http"
	"order-service/internal/api"
	"order-service/internal/rabbit"
)

func main() {

	err := rabbit.ConnectRabbitMQ("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("RabbitMQ connection error:", err)
	}
	defer rabbit.CloseRabbit()

	router := api.SetupRouter()

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", router)
}

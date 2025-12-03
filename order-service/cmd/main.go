package main

import (
	"fmt"
	"net/http"
	"order-service/internal/api"
)

func main() {
	router := api.SetupRouter()

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", router)
}

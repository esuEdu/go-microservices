package api

import (
	"net/http"
	"order-service/internal/api/handlers"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/orders", handlers.OrderHandler)

	return mux
}

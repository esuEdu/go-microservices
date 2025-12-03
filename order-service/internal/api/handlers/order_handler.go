package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/internal/model"
	"order-service/internal/rabbit"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createOrder(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	req := struct {
		UserId    string `json:"user_id"`
		ProductId string `json:"product_id"`
		Price     int    `json:"total_price"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	order := model.NewOrder(req.UserId, req.ProductId, req.Price)

	if err := rabbit.PublishOrder(order); err != nil {
		http.Error(w, "failed to publish order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "order created and sent to queue")

}

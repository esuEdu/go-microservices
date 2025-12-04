package processors

import (
	"encoding/json"
	"log"
	"payment-service/internal/models"
)

type PaymentPublisher interface {
	Publish(exchange, rk string, body []byte) error
}

type PaymentProcessor struct {
	pub PaymentPublisher
}

func NewPaymentProcessor(pub PaymentPublisher) *PaymentProcessor {
	return &PaymentProcessor{pub: pub}
}

// Implements rabbit.MessageHandler
func (p *PaymentProcessor) HandleMessage(body []byte) {
	var order struct {
		OrderID string  `json:"order_id"`
		Amount  float64 `json:"amount"`
	}

	if err := json.Unmarshal(body, &order); err != nil {
		log.Println("error:", err)
		return
	}

	payment := models.Payment{
		OrderID: order.OrderID,
		Amount:  order.Amount,
		Status:  "completed",
	}

	data, _ := json.Marshal(payment)
	p.pub.Publish("payment.exchange", "payment.completed", data)
}

package notifier

import (
	"encoding/json"
	"log"
)

type Payment struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}

func SendEmail(body []byte) {
	var p Payment
	if err := json.Unmarshal(body, &p); err != nil {
		log.Println("Failed to parse payment:", err)
		return
	}

	log.Printf("Email sent: Payment for order %s is %s\n", p.OrderID, p.Status)
}

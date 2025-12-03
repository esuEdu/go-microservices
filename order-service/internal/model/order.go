package model

import "github.com/google/uuid"

type Order struct {
	ID        uuid.UUID `json:"id"`
	UserId    string    `json:"user_id"`
	ProductId string    `json:"product_id"`
	Price     int       `json:"total_price"`
}

func NewOrder(userId, productId string, price int) Order {
	return Order{
		ID:        uuid.New(),
		UserId:    userId,
		ProductId: productId,
		Price:     price,
	}
}

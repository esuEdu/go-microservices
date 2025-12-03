package rabbit

import (
	"context"
	"encoding/json"
	"order-service/internal/model"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func PublishOrder(order model.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	_, err = RabbitChan.QueueDeclare(
		"orders_queue",
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return RabbitChan.PublishWithContext(
		ctx,
		"",             // exchange
		"orders_queue", // routing key
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConnection struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewRabbitConnection(url string) (*RabbitConnection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitConnection{Conn: conn, Ch: ch}, nil
}

func (r *RabbitConnection) Close() {
	if r.Ch != nil {
		r.Ch.Close()
	}
	if r.Conn != nil {
		r.Conn.Close()
	}
}

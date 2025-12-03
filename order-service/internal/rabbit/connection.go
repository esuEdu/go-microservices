package rabbit

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	RabbitConn *amqp.Connection
	RabbitChan *amqp.Channel
)

func ConnectRabbitMQ(url string) error {
	var err error

	RabbitConn, err = amqp.Dial(url)
	if err != nil {
		return err
	}

	RabbitChan, err = RabbitConn.Channel()
	if err != nil {
		return err
	}

	log.Println("Connected to RabbitMQ!")
	return nil
}

func CloseRabbit() {
	if RabbitChan != nil {
		RabbitChan.Close()
	}
	if RabbitConn != nil {
		RabbitConn.Close()
	}
}

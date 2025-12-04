package rabbit

type MessageHandler interface {
	HandleMessage([]byte)
}

type Consumer struct {
	handler MessageHandler
}

func NewConsumer(handler MessageHandler) *Consumer {
	return &Consumer{handler: handler}
}

func (c *Consumer) Consume(rc *RabbitConnection, queue string) error {
	msgs, err := rc.Ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			c.handler.HandleMessage(msg.Body)
		}
	}()

	return nil
}

package rabbit

type PaymentConsumer struct {
	notifier func([]byte)
}

func NewPaymentConsumer(n func([]byte)) *PaymentConsumer {
	return &PaymentConsumer{notifier: n}
}

func (c *PaymentConsumer) Consume(rc *RabbitConnection, queue string) error {
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
			c.notifier(msg.Body)
		}
	}()

	return nil
}

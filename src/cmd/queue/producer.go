package queue

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/suportebeloj/desafio-hitss/src/db/postgres"
)

type AmqpServiceProducer struct {
	ch *amqp.Channel
}

func NewAmqpServiceProducer(conn *amqp.Connection) *AmqpServiceProducer {
	service := &AmqpServiceProducer{}
	service.connect(conn)
	return service
}

func (c *AmqpServiceProducer) connect(conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	c.ch = ch
	return nil
}

func (c *AmqpServiceProducer) SendUser(ctx context.Context, queueName string, user postgres.CreateUserParams) error {
	defer c.ch.Close()
	q, err := c.ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = c.ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body})
	if err != nil {
		return err
	}

	return nil
}

package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type AmqpServiceConsumer struct {
	ch *amqp.Channel
	wg sync.WaitGroup
}

func NewAmqpServiceConsumer(conn *amqp.Connection) *AmqpServiceConsumer {
	service := &AmqpServiceConsumer{}
	if err := service.connect(conn); err != nil {
		log.Fatalln(err)
	}
	return service
}

func (c *AmqpServiceConsumer) connect(conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	c.ch = ch
	return nil
}

func (c *AmqpServiceConsumer) Consume(numWorkers int) error {
	defer c.ch.Close()
	q, err := c.ch.QueueDeclare(
		"newUser",
		false,

		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	msgs, err := c.ch.Consume(
		q.Name,
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

	for i := 0; i < numWorkers; i++ {
		c.wg.Add(1)
		go Worker(&c.wg, msgs)
	}
	c.wg.Wait()
	return nil
}

func Worker(wg *sync.WaitGroup, msgs <-chan amqp.Delivery) {
	defer wg.Done()

	for d := range msgs {
		body := string(d.Body)
		log.Println("Received Message: %s", body)
	}
}

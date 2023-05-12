package queue

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/suportebeloj/desafio-hitss/src/db/postgres"
	"github.com/suportebeloj/desafio-hitss/src/protocols"
	"log"
	"sync"
)

type AmqpServiceConsumer struct {
	ch                *amqp.Channel
	wg                sync.WaitGroup
	msgs              <-chan amqp.Delivery
	createUserService protocols.ICreateUserService
}

func NewAmqpServiceConsumer(conn *amqp.Connection, userService protocols.ICreateUserService) *AmqpServiceConsumer {
	service := &AmqpServiceConsumer{
		createUserService: userService,
	}
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

func (c *AmqpServiceConsumer) Consume(numWorkers int, queueName string) error {
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

	msgs, err := c.ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.msgs = msgs
	for i := 0; i < numWorkers; i++ {
		c.wg.Add(1)
		go c.Worker()
	}
	c.wg.Wait()
	return nil
}

func (c *AmqpServiceConsumer) Worker() {
	defer c.wg.Done()

	for d := range c.msgs {
		body := d.Body
		user := &postgres.CreateUserParams{}
		if err := json.Unmarshal(body, user); err != nil {
			log.Println("Failure to unmarshale message: ", err)
			return
		}
		fmt.Println(user)
		decodedUser, err := c.createUserService.ObfuscateInformation(context.Background(), *user, []string{"surname", "contact", "address", "cpf"}, 0)
		if err != nil {
			log.Println(err)
			return
		}

		if err := c.createUserService.Create(context.Background(), *decodedUser); err != nil {
			log.Println(err)
			return
		}
		d.Ack(false)
	}
}

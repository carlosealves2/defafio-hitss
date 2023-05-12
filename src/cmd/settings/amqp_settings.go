package settings

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

type AmqpConn struct {
	Conn *amqp.Connection
}

type amqpCredentials struct {
	dsn, user, pass, host, port string
}

var credential amqpCredentials

func (a *amqpCredentials) getDSn() string {
	if a.dsn != "" {
		return a.dsn
	}

	return fmt.Sprintf("amqp://%s:%s@%s:%s/", a.user, a.pass, a.host, a.port)
}

func NewAmqpConn() (*AmqpConn, error) {
	instance := &AmqpConn{}
	if err := instance.loadCredentials(); err != nil {
		return nil, err
	}

	if err := instance.connect(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (a *AmqpConn) loadCredentials() error {
	amqpDsn := os.Getenv("RABBITMQ_DSN")
	if amqpDsn != "" {
		credential.dsn = amqpDsn
		return nil
	}

	credential.user = os.Getenv("RABBITMQ_USER")
	credential.pass = os.Getenv("RABBITMQ_PASS")
	credential.host = os.Getenv("RABBITMQ_HOST")
	if credential.host == "" {
		credential.host = "localhost"
	}
	credential.port = os.Getenv("RABBITMQ_PORT")
	if credential.port == "" {
		credential.port = "5672"
	}

	return nil
}

func (a *AmqpConn) connect() error {
	conn, err := amqp.Dial(credential.getDSn())
	if err != nil {
		return err
	}
	a.Conn = conn
	return nil
}

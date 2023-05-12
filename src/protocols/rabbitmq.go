package protocols

import (
	"context"
	"github.com/suportebeloj/desafio-hitss/src/db/postgres"
)

type IRabbitMQProducerService interface {
	SendUser(ctx context.Context, queueName string, user postgres.CreateUserParams) error
}

type IRabbitMQConsumerService interface {
	Consume(numWorkers int, queueName string) error
}

package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/suportebeloj/desafio-hitss/src/cmd/queue"
	"github.com/suportebeloj/desafio-hitss/src/cmd/settings"
	"github.com/suportebeloj/desafio-hitss/src/db/postgres"
	"github.com/suportebeloj/desafio-hitss/src/usecases"
	"github.com/suportebeloj/desafio-hitss/src/utils/encrypter"
	"log"
	"time"
)

var (
	ctx context.Context = context.Background()
)

func init() {

}
func main() {
	defer settings.DbConn.Close()

	//queries := postgres.New(settings.DbConn)
	//userService := usecases.CreateUserService{}
	//apiService := api.NewFiberApiSerive(queries)
}

func Producer() {
	amqpConn, err := amqp.Dial("amqp://myuser:mypassword@localhost:5672/")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer amqpConn.Close()
	queueProducer := queue.NewAmqpServiceProducer(amqpConn)
	u := postgres.CreateUserParams{
		Name:    "carlos",
		Surname: "alves",
		Contact: "34234324",
		Address: "rua 2",
		Birth:   time.Now(),
		Cpf:     "123.233.233-23",
	}
	ctxTimou, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	_ = queueProducer.SendUser(ctxTimou, "newUser", u)
}

func Consumer() {
	amqpConn, err := amqp.Dial("amqp://myuser:mypassword@localhost:5672/")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer amqpConn.Close()

	queries := postgres.New(settings.DbConn)

	encrypterService := encrypter.NewEncrypter()
	userService := usecases.NewProcessUserData(queries, encrypterService)

	queueConsumer := queue.NewAmqpServiceConsumer(amqpConn, userService)
	queueConsumer.Consume(3, "newUser")
}

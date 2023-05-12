package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/suportebeloj/desafio-hitss/src/db/postgres"
	"github.com/suportebeloj/desafio-hitss/src/protocols"
	"log"
	"net/http"
	"strconv"
)

type FiberApiSerive struct {
	dbService   protocols.IDBService
	userService protocols.ICreateUserService
	producer    protocols.IRabbitMQProducerService
}

func NewFiberApiSerive(dbService protocols.IDBService, userService protocols.ICreateUserService, producer protocols.IRabbitMQProducerService) *FiberApiSerive {
	return &FiberApiSerive{dbService: dbService, userService: userService, producer: producer}
}

func (f FiberApiSerive) Run(addrs string) error {
	app := fiber.New()
	group := app.Group("/api/v1")
	group.Add("GET", "/user/:id", f.GetUser)
	group.Add("GET", "/users", f.ListUser)
	group.Add("POST", "/user/create", f.CreateUser)
	group.Add("DELETE", "/user/delete/:id", f.DeleteUser)
	group.Add("PUT", "/user/update/:id", f.UpdateUser)

	log.Fatalln(app.Listen(addrs))
	return nil
}

func (f FiberApiSerive) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	user, err := f.dbService.GetUser(context.Background(), int64(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("invalid ID")
	}

	return c.JSON(user)

}

func (f FiberApiSerive) ListUser(c *fiber.Ctx) error {
	result, err := f.dbService.ListUsers(context.Background())
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.JSON(result)
}

func (f FiberApiSerive) CreateUser(c *fiber.Ctx) error {
	createUserData := postgres.CreateUserParams{}

	if err := c.BodyParser(createUserData); err != nil {
		return err
	}

	if err := f.producer.SendUser(context.Background(), "new-user", createUserData); err != nil {
		return err
	}

	return c.SendStatus(http.StatusCreated)
}

func (f FiberApiSerive) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	updateUserParams := postgres.UpdateUserParams{ID: int64(id)}

	if err := c.BodyParser(updateUserParams); err != nil {
		return err
	}

	updated, err := f.dbService.UpdateUser(context.Background(), updateUserParams)
	if err != nil {
		return err
	}

	return c.JSON(updated)

}

func (f FiberApiSerive) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	deletedId, err := f.dbService.DeleteUser(context.Background(), int64(id))
	if err != nil {
		return err
	}

	return c.SendString(strconv.Itoa(int(deletedId)))
}

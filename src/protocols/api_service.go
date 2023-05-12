package protocols

import "github.com/gofiber/fiber/v2"

type IApiService interface {
	Run(addrs string) error
	GetUser(c *fiber.Ctx) error
	ListUser(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

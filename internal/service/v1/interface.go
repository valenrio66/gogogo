package v1

import "github.com/gofiber/fiber/v2"

type User interface {
	GetAllUser(ctx *fiber.Ctx) (err error)
}

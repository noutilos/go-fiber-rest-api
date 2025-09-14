package factory

import (
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}
package tasks

import (
	"net/http"

	"github.com/dev-khalid/go-fiber-rest-api/common/factory"
	"github.com/gofiber/fiber/v2"
)
type TaskController struct{}

func NewTaskController() factory.Controller {
	// Where do I put the DB instance?
	return &TaskController{}
}

// TODO: Should have a fixed return type which is Tasks struct, later we will add it to swagger docs as well.
func (c *TaskController) Create(ctx *fiber.Ctx) error {
	var task Task
	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	return ctx.JSON(fiber.Map{"message": "Task created", "task": task})
}

func (c *TaskController) Update(ctx *fiber.Ctx) error {
	return nil
}

func (c *TaskController) Delete(ctx *fiber.Ctx) error {
	return nil
}

func (c *TaskController) Get(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{"message": "Get task by ID"})
}

func (c *TaskController) GetAll(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "Get all tasks"})
}




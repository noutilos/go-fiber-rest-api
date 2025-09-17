package tasks

import (
	"github.com/dev-khalid/go-fiber-rest-api/common/utils"
	"github.com/gofiber/fiber/v2"
)

// singleton instance of the service
var controller = NewTaskController()

func SetupRoutes(app *fiber.App) {
	tasks := app.Group("/tasks")

	tasks.Get("/", controller.GetAll)
	tasks.Post("/", utils.CustomValidator(&Task{}), controller.Create)
	tasks.Get("/:id", controller.Get)
	tasks.Patch("/:id", controller.Update)
	tasks.Delete("/:id", controller.Delete)
}


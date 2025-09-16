package tasks

import (
	"net/http"
	"strconv"

	"github.com/dev-khalid/go-fiber-rest-api/common/factory"
	"github.com/gofiber/fiber/v2"
)
type TaskController struct{
	taskService *TaskService
}

func NewTaskController() factory.Controller {
	return &TaskController{
		taskService: NewTaskService(),
	}
}

func (c *TaskController) Create(ctx *fiber.Ctx) error {
	var task Task
	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	taskData, err := c.taskService.Create(&task)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task", "details": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Task created successfully", "task": taskData})
}

func (c *TaskController) Update(ctx *fiber.Ctx) error {
	return nil
}

func (c *TaskController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id") // Convert to int

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	err = c.taskService.Delete(taskID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete task", "details": err.Error(), "id": taskID})
	}

	return ctx.SendStatus(http.StatusNoContent)
}

func (c *TaskController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id") // Convert to int
	
	taskID, err := strconv.Atoi(id); 
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	task, err := c.taskService.Get(taskID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve task", "details": err.Error(), "id": taskID})
	}

	return ctx.JSON(fiber.Map{"message": "Get task by ID", "task": task})
}

func (c *TaskController) GetAll(ctx *fiber.Ctx) error {
	// Parse query parameters
	queryParams := NewTaskQueryParams()
	queryParams.ParseFromMap(ctx.Queries())

	tasks, err := c.taskService.GetAll(queryParams)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve tasks", "details": err.Error()})
	}

	// Return response with pagination info
	response := fiber.Map{
		"message": "Get all tasks",
		"tasks":   tasks,
		"pagination": fiber.Map{
			"page":  queryParams.Page,
			"limit": queryParams.Limit,
		},
	}

	return ctx.JSON(response)
}




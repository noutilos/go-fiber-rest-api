package main

import (
	"log"

	_ "github.com/dev-khalid/go-fiber-rest-api/config"
	"github.com/dev-khalid/go-fiber-rest-api/tasks"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	v1Api := app.Group("/v1/api")

	router := fiber.New(fiber.Config{
		Prefork: true,
	})
	tasks.SetupRoutes(router)

	v1Api.Mount("/", router)

	app.Listen(":8080")
	log.Println("Server is running on http://localhost:8080")
}
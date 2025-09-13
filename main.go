package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()


	app.Listen(":8080")
	log.Println("Server is running on http://localhost:8080")
}
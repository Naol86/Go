package main

import (
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!, naol",
	})
}

func main() {
	app := fiber.New()

	app.Get("/", hello)

	if err := app.Listen(":8000"); err != nil {
		panic(err)
	}
}
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Name string `json:"name"`
}

var book Book

func welcome(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	return c.JSON(book)
}

func main() {
	book.Name = "temp"
	api := fiber.New()

	api.Get("/api", welcome)

	log.Fatal(api.Listen(":8000"))

}
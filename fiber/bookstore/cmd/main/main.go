package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/naol86/Go/fiber/bookstore/pkg/route"
)
func main() {
	app := fiber.New()
	route.Routes(app)
	log.Fatal(app.Listen(":8000"))
}
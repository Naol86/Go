package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naol86/Go/fiber/bookstore/pkg/routes"
)

func main() {
	app := fiber.New()
	routes.Routes(app)
	app.Listen(":8000")

}
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naol86/Go/fiber/bookstore/pkg/models"
)


func Routes(app *fiber.App) {
	app.Get("/books", models.GetAllBooks)
}
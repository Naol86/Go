package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naol86/Go/fiber/bookstore/pkg/controllers"
)

func Routes(app *fiber.App) {
	app.Get("/api", controllers.Api)
	app.Get("/api/students", controllers.Students)
	app.Get("/api/students/:name", controllers.Student)
}

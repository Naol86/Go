package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naol86/Go/fiber/school/pkg/views"
)


func Routes(app *fiber.App){
	app.Get("/students", views.Students)
	app.Post("/students", views.CreateStudent)
	app.Get("/students/:id", views.Student)

}
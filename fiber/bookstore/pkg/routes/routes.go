package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naol86/Go/fiber/bookstore/pkg/models"
)


func Routes(app *fiber.App) {
	// school routes
	app.Post("/schools", models.CreateSchools)
	app.Get("/schools", models.GetAllSchools)
	app.Delete("/schools/:school_id", models.DeleteSchool)
	// app.Put("/schools", models.UpdateSchool)

	// department routes
	app.Post("/department/:school_id", models.CreateDepartment)
	app.Get("/department/:school_id", models.GetDepartments)
	app.Delete("/department/:department_id", models.DeleteDepartment)

	// course routes
	app.Post("/course/:department_id", models.CreateCourse)
	app.Get("/course/:department_id", models.GetAllCourse)
	app.Delete("/course/:course_id", models.DeleteCourse)

	// book routes
	app.Post("/books/:course_id", models.CreateBook)
	app.Get("/books/:course_id", models.GetAllBooks)
	app.Get("/books", models.GetBooks)
	app.Delete("/books/:book_id", models.DeleteBook)

}
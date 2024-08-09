package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naol86/mon-go/controller"
)



func Routes(app *fiber.App) {
	api := app.Group("/api");
	api.Get("/", controller.Hello);
	api.Get("/users", controller.GetAllUsers);
	api.Post("/users", controller.CreateUser);
	api.Post("/login", controller.LoginUser);

}
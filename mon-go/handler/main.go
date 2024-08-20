package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naol86/mon-go/controller"
)



func Routes(app *fiber.App) {
	api := app.Group("/api");
	api.Get("/", controller.Hello);
	api.Get("/users", func  (c *fiber.Ctx) error {
		fmt.Println("here we go again ")
		return c.Next()} ,controller.GetAllUsers);
	api.Post("/users", controller.CreateUser);
	api.Post("/signin", controller.SigninUser);
	api.Post("/check", controller.Check)

}
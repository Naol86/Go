package main

import (
	"github.com/gofiber/fiber/v2"
	routes "github.com/naol86/mon-go/handler"
)


func main() {
	// initialize db connection
	// db.ConnectToMongoDB();


	app := fiber.New();


	// Routes
	routes.Routes(app);

	app.Listen(":3000")
}
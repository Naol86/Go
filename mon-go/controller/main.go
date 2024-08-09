package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naol86/mon-go/database"
)

var db = database.GetDB();
var collection = db.Database("testMongo").Collection("users");

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!");
}


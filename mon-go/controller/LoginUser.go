package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naol86/mon-go/models"
	"github.com/naol86/mon-go/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginUser(c *fiber.Ctx) error {
	var user models.User
	var credentials models.UserLoginForm;
	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid JSON format",
		})
	}

	filter := bson.M{"email": credentials.Email};
	err := collection.FindOne(c.Context(), filter).Decode(&user);
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
			"success": false,
		})
	}

	if err := utils.CheckPasswordHash(credentials.Password, user.Password);  !err {
		return c.Status(401).JSON(fiber.Map{
			"message": "wrong password",
			"success": false,
		})
	}


	fmt.Println(user);
	return c.JSON(fiber.Map{
		"message": "User found",
	})
}
package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naol86/mon-go/models"
	"github.com/naol86/mon-go/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.UserRegisterForm;
	err := c.BodyParser(&user);
	if err != nil {
		return c.SendString("Error parsing user");
	}

	if user.Password != user.ConfirmPassword {
		return c.SendString("Password and confirm password do not match");
	}

	hash, _ := utils.HashPassword(user.Password);

	var newUser models.User;
	newUser.Name = user.Name;
	newUser.Email = user.Email;
	newUser.Password = hash;

	Id,_ := collection.InsertOne(c.Context(), newUser);

	// newUser,_ := collection.InsertOne(c.Context(), user);
	// user.ID = newUser.InsertedID.(primitive.ObjectID);

	response := fiber.Map{
		"message": "User created successfully",
		"success": true,
		"user": fiber.Map{
			"id": Id.InsertedID.(primitive.ObjectID),
			"name": user.Name,
			"email": user.Email,
		},
		"errors": nil,
		"count": 0,
	}


	return c.JSON(response);
}
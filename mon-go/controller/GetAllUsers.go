package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/naol86/mon-go/models"
	"go.mongodb.org/mongo-driver/bson"
)


func GetAllUsers(c *fiber.Ctx) error {

	// cursor, err := collection.Find(c.Context(), nil);
	cursor, err := collection.Find(context.Background(), bson.D{});
	if err != nil {
		return c.Status(500).SendString(err.Error());
	}
	defer cursor.Close(context.Background());

	var users []models.User;
	for cursor.Next(context.Background()) {
		var user models.User;
		if err := cursor.Decode(&user); err != nil {
			return c.Status(500).SendString(err.Error());
		}
		users = append(users, user);
	}

	return c.JSON(users);

}
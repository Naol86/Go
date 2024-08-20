package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naol86/mon-go/auth"
)



func Check(c *fiber.Ctx) error {
	token := c.Get("Authorization");

	claims,err := auth.ValidateToken(token[len("Bearer "):]);
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid token",
			"success": false,
		})
	}

	fmt.Println(claims);
	
	return c.JSON(fiber.Map{
		"message": "Valid token",
		"success": true,
	});

}
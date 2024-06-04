package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naol86/Go/fiber/bookstore/pkg/models"
)

func Api(c *fiber.Ctx) error {
	return c.JSON("hello naol")
}

func Students(c *fiber.Ctx) error {
	var students []models.Student
	students = models.GetAll()
	return c.JSON(students)
}

func Student(c *fiber.Ctx) error {
	var student models.Student
	student = *models.GetByName(c.Params("name"))
	return c.JSON(student)
}

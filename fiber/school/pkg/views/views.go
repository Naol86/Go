package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naol86/Go/fiber/school/pkg/models"
)

func Students(c *fiber.Ctx) error{
	students := models.GetAllStudent()
	return c.JSON(students)
}

func Student(c *fiber.Ctx) error {
	student := models.GetStudentById(c.Params("id"))
	return c.JSON(student)
}

func CreateStudent(c *fiber.Ctx) error {
	student := new(models.Student)
	if err := c.BodyParser(student); err != nil{
		return c.Status(503).SendString(err.Error())
	}
	student.CreateStudent()
	return c.JSON(student)
}


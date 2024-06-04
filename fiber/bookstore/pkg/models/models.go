package models

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/naol86/Go/fiber/bookstore/pkg/database"
	"gorm.io/gorm"
)

// School represents the School table
type School struct {
	gorm.Model
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Departments []Department `gorm:"foreignKey:SchoolID" json:"departments,omitempty"`
}

// Department represents the Department table
type Department struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SchoolID    uint    `json:"school_id"`
	Courses     []Course `gorm:"foreignKey:DepartmentID" json:"courses,omitempty"`
}

// Course represents the Course table
type Course struct {
	gorm.Model
	Name         string `json:"name"`
	Description  string `json:"description"`
	Code         string `json:"code"`
	DepartmentID uint   `json:"department_id"`
	Books        []Books `gorm:"foreignKey:CourseID" json:"books,omitempty"`
}

// Book represents the Book table
type Books struct {
	gorm.Model
	Name         string `json:"name"`
	Description  string `json:"description"`
	File         string `json:"file"`
	Image        string `json:"image"`
	CourseID     uint   `json:"course_id"`
}


var DB *gorm.DB

func init() {
	database.Connect()
	DB = database.DB
	err := DB.AutoMigrate(&School{}, &Department{}, &Course{}, &Books{})
	if err != nil {
		log.Fatal("fail to connect with database")
	}
}

func GetAllBooks(c *fiber.Ctx) error {
	return c.SendString("hello world")
}
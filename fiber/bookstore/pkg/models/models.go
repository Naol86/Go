package models

import (
	"log"
	"strconv"

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

func (s *School) Save() {
	DB.Create(&s)
}


// Department represents the Department table
type Department struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SchoolID    uint    `json:"school_id"`
	Courses     []Course `gorm:"foreignKey:DepartmentID" json:"courses,omitempty"`
}

func (d *Department) Save() {
	DB.Create(&d)
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

func (c *Course) Save() {
	DB.Create(&c)
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

func (b *Books) Save() {
	DB.Create(&b)
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

// start school
func CreateSchools(c *fiber.Ctx) error {
	school := new(School)
	err := c.BodyParser(school)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	school.Save()
	return c.JSON(school)
}

func GetAllSchools(c *fiber.Ctx) error {
	var schools []School
	DB.Find(&schools)
	return c.JSON(schools)
}


// end school

// start department

func CreateDepartment(c *fiber.Ctx) error {
	department := new(Department)
	err := c.BodyParser(&department)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	School_ID := c.Params("school_id")
	SchoolID, err := strconv.ParseUint(School_ID, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid school ID")
	}
	department.SchoolID = uint(SchoolID)
	department.Save()
	return c.JSON(department)
}

func GetDepartments(c *fiber.Ctx) error {
	Did := c.Params("school_id")
	d_id, err := strconv.ParseUint(Did, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid school ID")
	}
	var departments []Department
	DB.Where("school_id = ?", d_id).Find(&departments)
	return c.JSON(departments)
}

// end department

// start course
func CreateCourse(c *fiber.Ctx) error {
	course := new(Course)
	err := c.BodyParser(&course)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	Did := c.Params("department_id")
	d_id, err := strconv.ParseUint(Did, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid department")
	}
	course.DepartmentID = uint(d_id)
	course.Save()
	return c.JSON(course)

}

func GetAllCourse(c *fiber.Ctx) error {
	Cid := c.Params("department_id")
	c_id, err := strconv.ParseUint(Cid, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid department")
	}
	var courses []Course
	DB.Where("department_id = ?", c_id).Find(&courses)
	return c.JSON(courses)
}
// end course



// books

func CreateBook(c *fiber.Ctx) error {
	book := new(Books)
	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	CId := c.Params("course_id")
	c_id, err := strconv.ParseUint(CId, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid course")
	}
	book.CourseID = uint(c_id)
	book.Save()
	return c.JSON(book)
}

func GetAllBooks(c *fiber.Ctx) error {
	Cid := c.Params("course_id")
	c_id, err := strconv.ParseUint(Cid, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid course")
	}
	var books []Books
	DB.Where("course_id = ?", c_id).Find(&books)
	return c.JSON(books)
}

// end books
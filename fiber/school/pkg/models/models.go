package models

import (
	// "github.com/jinzhu/gorm"
	"github.com/naol86/Go/fiber/school/pkg/database"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name       string `json:"name"`
	Age        string `json:"age"`
	Sex        string `json:"sex"`
	Department string `json:"department"`
}

var DB *gorm.DB

func init() {
	database.Connect()
	DB = database.GetDB()
	
	DB.AutoMigrate(&Student{})
}

func GetAllStudent() []Student {
	var students []Student
	DB.Find(&students)
	return students
}

func GetStudentById(id string) Student {
	var student Student
	DB.First(&student, id)
	return student
}

func (s *Student) CreateStudent() *Student{
	DB.Create(&s)
	return s
}

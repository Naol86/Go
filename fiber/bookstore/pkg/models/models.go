package models

type Student struct {
	Name string `json:"name"`
	Age string `json:"age"`
	Department string `json:"department"`
}

var Students []Student

func init() {
	Students = append(Students , Student{
	Name: "naol",
	Age: "21",
	Department: "software",
	})
	Students = append(Students , Student{
		Name: "temp",
		Age: "21",
		Department: "software",
		})
}

func GetAll() []Student {
	return Students
}

func GetByName(name string) *Student{
	for _, item := range Students {
		if item.Name == name {
			return &item
		}
	}
	return nil
}
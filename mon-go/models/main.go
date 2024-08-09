package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Email  string             `bson:"email,omitempty"`
	Password string					 `bson:"password,omitempty"`
}

type UserRegisterForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserLoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string             `bson:"name,omitempty" json:"name,omitempty"`
	Username string					 `bson:"username,omitempty" json:"username,omitempty"`
	Email  string             `bson:"email,omitempty" json:"email,omitempty"` 
	Password string					 `bson:"password,omitempty" json:"-"`
}

type UserRegisterForm struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserLoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
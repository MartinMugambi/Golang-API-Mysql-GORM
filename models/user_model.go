package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

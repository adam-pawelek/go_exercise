package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string
	Password string
}

type Body struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

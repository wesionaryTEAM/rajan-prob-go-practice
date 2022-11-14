package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body string
	Amount int
	UserID uint
	User User
}
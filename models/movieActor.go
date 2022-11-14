package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name   string
	Actors []Actor `gorm:"many2many:movie_actors;"`
}

type Actor struct {
	gorm.Model
	Name string
	Movies []Movie `gorm:"many2many:movie_actors;"`
  }
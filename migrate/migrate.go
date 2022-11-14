package main

import (
	initialize "example/rajan-prob-go-practice/Initialize"
	"example/rajan-prob-go-practice/models"
)

func init() {
	initialize.LoadEnv()
	initialize.ConnectDB()
}
func main() {
initialize.DB.AutoMigrate(&models.Post{})
initialize.DB.AutoMigrate(&models.User{})
initialize.DB.AutoMigrate(&models.Movie{})
initialize.DB.AutoMigrate(&models.Actor{})
}

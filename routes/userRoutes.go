package routes

import (
	"example/rajan-prob-go-practice/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser)
	router.GET("/user", controllers.GetUsers)
}
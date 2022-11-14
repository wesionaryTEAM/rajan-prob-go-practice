package routes

import (
	"example/rajan-prob-go-practice/controllers"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(router *gin.Engine) {
	router.POST("/movie", controllers.CreateMovie)
	router.GET("/movie", controllers.GetMovies)
	router.POST("/actor", controllers.CreateActor)
	router.GET("/actor", controllers.GetActors)

}
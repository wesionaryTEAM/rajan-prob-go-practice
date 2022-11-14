package controllers

import (
	initialize "example/rajan-prob-go-practice/Initialize"
	"example/rajan-prob-go-practice/models"
	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {

var movie models.Movie

c.Bind(&movie)

result := initialize.DB.Create(&movie) 


if result.Error!=nil{
	c.Status(400)
	return
}
	
	c.JSON(200, gin.H{
		"movie":movie,
	})
}

func GetMovies(c *gin.Context){

	// Get all records
	var movies []models.Movie

   result := initialize.DB.Preload("Actors").Find(&movies)

   
if result.Error!=nil{
	c.Status(400)
	return
}
c.JSON(200, gin.H{
	"movies":movies,
})
}

func CreateActor(c *gin.Context) {

	var actor models.Actor
	
	c.Bind(&actor)
	
	result := initialize.DB.Create(&actor) 
	
	
	if result.Error!=nil{
		c.Status(400)
		return
	}
		
		c.JSON(200, gin.H{
			"actor":actor,
		})
	}
	
	func GetActors(c *gin.Context){
	
		// Get all records
		var actors []models.Actor
	
	   result := initialize.DB.Preload("Movies").Find(&actors)
	
	   
	if result.Error!=nil{
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"actors":actors,
	})
	}
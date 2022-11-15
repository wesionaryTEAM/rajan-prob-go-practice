package controllers

import (
	initialize "example/rajan-prob-go-practice/Initialize"
	"example/rajan-prob-go-practice/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

var user models.User

if err:=c.ShouldBindJSON(&user); err!=nil{
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}


result := initialize.DB.Create(&user) 


if result.Error!=nil{
	c.Status(400)
	return
}
	
	c.JSON(200, gin.H{
		"user":user,
	})
}

func GetUsers(c *gin.Context){

	// Get all records
	var users []models.User

   result := initialize.DB.Find(&users)

   //Selecting a posts with amount > avg of all amount
    // subQuery:=initialize.DB.Table("posts").Select("AVG(amount)")
    // result := initialize.DB.Where("amount > (?)", subQuery).Find(&posts)
	

	//selecting multiples intable
	//result := initialize.DB.Where("(title, amount) IN ?", [][]interface{}{{"test4", 4}, {"test3", 2}}).Find(&posts)
if result.Error!=nil{
	c.Status(400)
	return
}
c.JSON(200, gin.H{
	"users":users,
})
}
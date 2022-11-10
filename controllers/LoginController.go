package controllers

import (
	"example/rajan-prob-go-practice/models"
	"example/rajan-prob-go-practice/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService service.JWTService
}

func Loginhandler(loginService service.LoginService,jwtService service.JWTService) LoginController{
	return &loginController{	
		loginService: loginService,
		jwtService: jwtService,
	}
}

func (controller *loginController) 	Login(ctx *gin.Context) string{
 var loginData models.Login
 err:=ctx.ShouldBind(&loginData)
 if err !=nil{
	return "no data Found"

 }
 isUserAuthenticated:=controller.loginService.LoginUser(loginData.Email,loginData.Password)

 if isUserAuthenticated{
	return controller.jwtService.GenerateToken(loginData.Email,isUserAuthenticated)
 }
 return ""
}
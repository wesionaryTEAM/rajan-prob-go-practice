package main

import (
	initialize "example/rajan-prob-go-practice/Initialize"
	"example/rajan-prob-go-practice/controllers"
	"example/rajan-prob-go-practice/routes"
	"example/rajan-prob-go-practice/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init(){
	initialize.LoadEnv()
	initialize.ConnectDB()
}
func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controllers.LoginController = controllers.Loginhandler(loginService, jwtService)

	server := gin.New()

	routes.PostRoutes(server);
	routes.UserRoutes(server);
	routes.MovieRoutes(server);
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	port := "3000"
	server.Run(":" + port)
}
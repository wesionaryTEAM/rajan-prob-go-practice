package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/public", "./public")
	router.StaticFS("/more_static", http.Dir("public"))
	//router.StaticFile("/public/favicon.ico", "./public/3d-lampa-rozetka-ideya-elektrichestvo-155370.jpeg")
	// router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))




	//serving data from file
	router.GET("/local/file", func(c *gin.Context) {
		c.File("public/file.go")
	  })
	
//serving from reader
	  router.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
		  c.Status(http.StatusServiceUnavailable)
		  return
		}
	
		reader := response.Body
		 defer reader.Close()
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
	
		extraHeaders := map[string]string{
		  "Content-Disposition": `attachment; filename="gopher.png"`,
		}
	
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader,extraHeaders)
	  })
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
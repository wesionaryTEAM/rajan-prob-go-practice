package main

// import (
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// type Person struct {
//         Name       string    `json:"name"`
//         Address    string    `json:"address"`
//         Birthday   time.Time `json:"birthday" time_format:"2006-01-02" time_utc:"1"`
// }

// func main() {
//   route := gin.Default()
//   route.POST("/testing", startPage)
//   route.Run(":8085")
// }

// func startPage(c *gin.Context) {
//   var person Person
//   // If `GET`, only `Form` binding engine (`query`) used.
//   // If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
//   // See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
//         if c.ShouldBind(&person) == nil {
//                 log.Println(person.Name)
//                 log.Println(person.Address)
//                 log.Println(person.Birthday)

//         }

//   c.String(http.StatusOK, "Success")
// }
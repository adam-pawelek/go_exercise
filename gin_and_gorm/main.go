package main

import (
	"fmt"

	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/controllers"
	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", controllers.Signup)
	fmt.Println("sadf")
	r.Run() // listen and serve on 0.0.0.0:8080
}
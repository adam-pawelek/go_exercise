package main

import (
	"fmt"

	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/controllers"
	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/initializers"
	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/middleware"
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
	r.POST("login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	fmt.Println("sadf")
	r.Run() // listen and serve on 0.0.0.0:8080
}

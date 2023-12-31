package main

import (
	"fmt"
	"os"

	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/controllers"
	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/docs"
	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/initializers"
	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	//initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apiKey JWT
// @type     apiKey
// @in       cookie
// @name     Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("PORT"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}

package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	docs "simple_api/docs"
	"simple_api/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/weapons", v1.GetWeapons)
	r.GET("/weapon", v1.GetWeapon)
	r.POST("/weapon", v1.PostWeapon)
	r.DELETE("/weapon/:type", v1.DeleteWeapon)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	return r
}

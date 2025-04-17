package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpnt17/kumo-classroom-be/internal/controllers"
	"github.com/tuanpnt17/kumo-classroom-be/internal/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	uc := controllers.NewUserController()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/users/:id", uc.GetUserById)
		v1.GET("/users", uc.GetAll)
	}

	v2 := r.Group("/api/v2")
	{
		v2.Use(middlewares.AuthMiddleware())
		v2.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong with v2 api",
			})
		})
	}
	return r
}

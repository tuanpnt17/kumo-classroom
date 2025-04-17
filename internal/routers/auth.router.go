package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpnt17/kumo-classroom-be/internal/middlewares"
)

type AuthRouter struct {
}

func (a *AuthRouter) InitAuthRouter(router *gin.RouterGroup) {
	publicAuthRouter := router.Group("/auth")
	{
		publicAuthRouter.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
		publicAuthRouter.POST("/register", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
		publicAuthRouter.POST("/refresh", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
	}

	privateAuthRouter := router.Group("/auth")
	{
		privateAuthRouter.Use(middlewares.Authentication())
		//privateAuthRouter.Use(middlewares.Authorization())
		privateAuthRouter.POST("/logout", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
	}
}

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpnt17/kumo-classroom-be/internal/middlewares"
	"github.com/tuanpnt17/kumo-classroom-be/internal/wire"
)

type AuthRouter struct {
}

func (a *AuthRouter) InitAuthRouter(router *gin.RouterGroup) {
	ac := wire.InitAuthController()
	publicAuthRouter := router.Group("/auth")
	{
		publicAuthRouter.POST("/login", ac.Login)
		publicAuthRouter.POST("/register", ac.Register)
		publicAuthRouter.POST("/refresh", ac.Refresh)
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

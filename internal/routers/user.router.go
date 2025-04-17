package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpnt17/kumo-classroom-be/internal/middlewares"
	"github.com/tuanpnt17/kumo-classroom-be/internal/wire"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(router *gin.RouterGroup) {

	userController := wire.InitUserController()

	publicUserRouter := router.Group("/users")
	{
		publicUserRouter.GET("", userController.GetAll)
		publicUserRouter.GET("/info", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})

	}

	privateUserRouter := router.Group("/user")
	{
		privateUserRouter.Use(middlewares.Authentication())
		//privateUserRouter.Use(middlewares.Authorization())

	}
}

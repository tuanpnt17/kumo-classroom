package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpnt17/kumo-classroom-be/global"
	"github.com/tuanpnt17/kumo-classroom-be/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == gin.DebugMode {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middleware
	//r.Use() // logging
	//r.Use() // cors

	userRouter := routers.RouterGroupApp.User
	authRouter := routers.RouterGroupApp.Auth

	// router group
	v1 := r.Group("/api/v1")
	{
		userRouter.InitUserRouter(v1)
		authRouter.InitAuthRouter(v1)
	}

	return r
}

package initialize

import (
	"github.com/tuanpnt17/kumo-classroom-be/global"
	"strconv"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitPostgres()
	// InitRedis()
	r := InitRouter()

	global.Logger.Debug("All initializations are done")
	if err := r.Run(":" + strconv.Itoa(global.Config.Server.Port)); err != nil {
		global.Logger.Error("Failed to start server: " + err.Error())
		panic(err)
	}
	global.Logger.Debug("Server started on port " + strconv.Itoa(global.Config.Server.Port))
}

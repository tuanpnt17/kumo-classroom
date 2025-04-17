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

	global.Logger.Debug("All initializations are done")
	r := InitRouter()

	if err := r.Run(":" + strconv.Itoa(global.Config.Server.Port)); err != nil {
		panic(err)
	}
}

package initialize

import (
	"github.com/tuanpnt17/kumo-classroom-be/global"
	"github.com/tuanpnt17/kumo-classroom-be/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewZapLogger(global.Config.ZapLogger)
}

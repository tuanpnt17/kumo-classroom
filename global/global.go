package global

import (
	"github.com/tuanpnt17/kumo-classroom-be/pkg/logger"
	"github.com/tuanpnt17/kumo-classroom-be/pkg/settings"
	"gorm.io/gorm"
)

var (
	Config *settings.AppConfig
	Logger *logger.ZapLogger

	AppDB *gorm.DB
)

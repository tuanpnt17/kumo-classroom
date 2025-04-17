package global

import (
	"github.com/tuanpnt17/kumo-classroom-be/pkg/logger"
	"github.com/tuanpnt17/kumo-classroom-be/pkg/settings"
)

var (
	Config *settings.AppConfig
	Logger *logger.ZapLogger
)

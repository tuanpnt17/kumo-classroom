//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/tuanpnt17/kumo-classroom-be/internal/controllers"
)

func InitAuthController() *controllers.AuthController {
	wire.Build(ControllerSet)
	return new(controllers.AuthController)
}

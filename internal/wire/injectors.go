//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/tuanpnt17/kumo-classroom-be/internal/controllers"
)

func InitUserController() *controllers.UserController {
	wire.Build(ControllerSet)
	return new(controllers.UserController)
}

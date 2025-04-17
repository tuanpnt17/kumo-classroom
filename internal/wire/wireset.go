package wire

import (
	"github.com/google/wire"
	"github.com/tuanpnt17/kumo-classroom-be/internal/controllers"
	"github.com/tuanpnt17/kumo-classroom-be/internal/repositories"
	"github.com/tuanpnt17/kumo-classroom-be/internal/services"
)

var RepoSet = wire.NewSet(
	repositories.NewUserRepository,
)

var ServiceSet = wire.NewSet(
	RepoSet,
	services.NewUserService,
)

var ControllerSet = wire.NewSet(
	ServiceSet,
	controllers.NewUserController,
)

package wire

import (
	"github.com/google/wire"
	"github.com/tuanpnt17/kumo-classroom-be/internal/controllers"
	"github.com/tuanpnt17/kumo-classroom-be/internal/repositories"
	"github.com/tuanpnt17/kumo-classroom-be/internal/usecases"
)

var RepoSet = wire.NewSet(
	repositories.NewUnitOfWork,
	repositories.NewUserRepository,
)

var UsecaseSet = wire.NewSet(
	RepoSet,
	usecases.NewLoginUsecase,
)

var ControllerSet = wire.NewSet(
	UsecaseSet,
	controllers.NewAuthController,
)

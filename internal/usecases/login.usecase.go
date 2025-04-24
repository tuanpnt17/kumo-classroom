package usecases

import (
	"context"
	"github.com/tuanpnt17/kumo-classroom-be/internal/domain/entities"
	"github.com/tuanpnt17/kumo-classroom-be/internal/domain/models"
)

type loginUsecase struct {
	userRepository entities.IUserRepository
}

func (l *loginUsecase) HasUser(ctx context.Context, email string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (l *loginUsecase) NewJwtToken(user *entities.User) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (l *loginUsecase) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewLoginUsecase(ur entities.IUserRepository) models.ILoginUsecase {
	return &loginUsecase{
		userRepository: ur,
	}
}

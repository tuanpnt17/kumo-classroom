package repositories

import (
	"context"
	"github.com/tuanpnt17/kumo-classroom-be/global"
	"github.com/tuanpnt17/kumo-classroom-be/internal/domain/entities"
	"github.com/tuanpnt17/kumo-classroom-be/internal/domain/interfaces"
	"time"
)

type userRepository struct {
	unitOfWork interfaces.IUnitOfWork
}

func (ur userRepository) CheckExistByEmail(ctx context.Context, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(global.Config.Postgres.TimeoutSeconds)*time.Second)
	defer cancel()

	count, err := ur.unitOfWork.CountRows(ctx, &entities.User{Email: email})
	return count > 0, err
}

func NewUserRepository(unitOfWork interfaces.IUnitOfWork) entities.IUserRepository {
	return &userRepository{
		unitOfWork: unitOfWork,
	}
}

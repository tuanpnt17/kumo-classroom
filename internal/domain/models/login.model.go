package models

import (
	"context"
	"github.com/tuanpnt17/kumo-classroom-be/internal/domain/entities"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ILoginUsecase interface {
	HasUser(ctx context.Context, email string) (bool, error)
	NewJwtToken(user *entities.User) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
}

package services

import "github.com/tuanpnt17/kumo-classroom-be/internal/repositories"

type IUserService interface {
	GetAllUsers() []string
}
type userService struct {
	userRepository *repositories.IUserRepository
}

func (us *userService) GetAllUsers() []string {
	return (*us.userRepository).GetAll()
}

func NewUserService(ur repositories.IUserRepository) IUserService {
	return &userService{
		userRepository: &ur,
	}
}

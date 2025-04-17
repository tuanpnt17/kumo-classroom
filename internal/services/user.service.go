package services

import "github.com/tuanpnt17/kumo-classroom-be/internal/repositories"

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(),
	}
}

func (us *UserService) GetAllUsers() []string {
	return us.userRepository.GetAll()
}

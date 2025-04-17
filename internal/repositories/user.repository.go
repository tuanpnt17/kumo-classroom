package repositories

type IUserRepository interface {
	GetAll() []string
}

type userRepository struct{}

func (ur *userRepository) GetAll() []string {
	return []string{"user1", "user2", "user3"}
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

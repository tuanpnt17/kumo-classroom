package repositories

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetAll() []string {
	users := []string{"Tuan", "Anh", "Quynh"}
	return users
}

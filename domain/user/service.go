package user

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) CreateUser(user *User) error {
	return us.userRepository.Create(user)
}

func (us *UserService) GetUserByID(id int) (*User, error) {
	return us.userRepository.GetByID(id)
}

func (us *UserService) UpdateUser(user *User) error {
	return us.userRepository.Update(user)
}

func (us *UserService) DeleteUser(id int) error {
	return us.userRepository.Delete(id)
}

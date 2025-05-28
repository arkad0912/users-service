package userService

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *User) (*User, error) {
	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) GetUserByID(id uint) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(id uint, user *User) (*User, error) {
	err := s.repo.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	return s.repo.GetUserByID(id)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

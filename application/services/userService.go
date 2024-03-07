package services

import (
	"fmt"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/application/services/auth"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/repository"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type UserService struct {
	userRepository repository.UserRepository
	authService    auth.AuthService
}

func CreateUserService(userRepository repository.UserRepository, authService auth.AuthService) *UserService {
	return &UserService{
		userRepository: userRepository,
		authService:    authService,
	}
}

func (s *UserService) NewUser(user restModel.UserRestModel) (string, error) {
	email, role, err := s.userRepository.NewUser(user)
	fmt.Println(email)
	if err != nil {
		return "", err
	}
	token, err := s.authService.GenerateToken(user.Email, role)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (s *UserService) AuthenticateUser(email, password string) (string, *entity.User, error) {
	user, err := s.userRepository.AuthenticateUser(email, password)
	if err != nil {
		return "", nil, fmt.Errorf(err.Error())
	}
	token, err := s.authService.GenerateToken(email, user.Role)
	if err != nil {
		return "", nil, fmt.Errorf(err.Error())
	}
	return token, user, nil
}
func (s *UserService) GetUser(email string) (*entity.User, error) {
	return s.userRepository.GetUserByEmail(email)
}
func (s *UserService) GetAllUser() ([]*entity.User, error) {
	return s.userRepository.GetAllUser()
}
func (s *UserService) DeleteUser(email string) error {
	return s.userRepository.DeleteUser(email)
}
func (s *UserService) EditUser(user entity.User) (*entity.User, error) {
	return s.userRepository.EditUser(user)
}

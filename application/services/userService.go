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
	email, err := s.userRepository.NewUser(user)
	fmt.Println(email)
	if err != nil {
		return "", err
	}
	token, err := s.authService.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (s *UserService) AuthenticateUser(email, password string) (string, error) {
	_, err := s.userRepository.AuthenticateUser(email, password)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	token, err := s.authService.GenerateToken(email)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return token, nil
}
func (s *UserService) GetUser(email string) (*entity.User, error) {
	return s.userRepository.GetUserByEmail(email)
}
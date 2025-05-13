package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"url/internal/user"
	"url/pkg/di"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	ExistedUser, _ := service.UserRepository.FindUserByEmail(email)
	if ExistedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	NewUser := user.User{
		Email:    email,
		Password: string(HashedPassword),
		Name:     name,
	}
	_, err = service.UserRepository.Create(&NewUser)
	if err != nil {
		return "", err
	}
	return NewUser.Email, nil
}
func (service *AuthService) Login(email, password string) (string, error) {
	ExistedUser, _ := service.UserRepository.FindUserByEmail(email)
	if ExistedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}
	err := bcrypt.CompareHashAndPassword([]byte(ExistedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}
	return ExistedUser.Email, nil
}

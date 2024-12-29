package services

import (
	"errors"
	"news-portal/internal/models/user"
	"news-portal/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *user.User) error
	GetUserByID(id uint) (*user.User, error)
	GetUserByEmail(email string) (*user.User, error)
	UpdateUser(user *user.User) error
	DeleteUser(id uint) error
	VerifyPassword(user *user.User, password string) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(user *user.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepository.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (*user.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *userService) GetUserByEmail(email string) (*user.User, error) {
	return s.userRepository.GetUserByEmail(email)
}

func (s *userService) UpdateUser(user *user.User) error {
	return s.userRepository.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepository.DeleteUser(id)
}

func (s *userService) VerifyPassword(user *user.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("invalid login credentials")
	}
	return nil
}

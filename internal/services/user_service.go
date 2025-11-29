package services

import (
	"errors"

	"github.com/aswearingen91/account-service/internal/models"
	"github.com/aswearingen91/account-service/internal/repositories"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(username string) (*models.User, error)
	GetUser(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

type userService struct {
	users repositories.UserRepository
}

func NewUserService(users repositories.UserRepository) UserService {
	return &userService{users}
}

func (s *userService) CreateUser(username string) (*models.User, error) {
	_, err := s.users.GetByUsername(username)
	if err == nil {
		// found an existing user
		return nil, errors.New("username already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// unexpected DB error
		return nil, err
	}

	user := &models.User{
		Username: username,
	}

	if err := s.users.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	return s.users.GetByID(id)
}

func (s *userService) GetUserByUsername(username string) (*models.User, error) {
	return s.users.GetByUsername(username)
}

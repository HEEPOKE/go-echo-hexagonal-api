package services

import (
	"github.com/HEEPOKE/go-gin-hexagonal-api/internal/domains/models"
	"github.com/HEEPOKE/go-gin-hexagonal-api/internal/interfaces"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	return s.userRepository.GetAll()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepository.Create(user)
}
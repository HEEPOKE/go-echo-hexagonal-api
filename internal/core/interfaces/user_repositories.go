package interfaces

import "github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"

type UserRepository interface {
	GetAll() ([]*models.User, error)
	GetByID(id int) (*models.User, error)
	Create(user *models.User) error
}

package interfaces

import "github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"

type AuthRepository interface {
	Login(email, password string) (*models.User, error)
	Register(user *models.User) error
}

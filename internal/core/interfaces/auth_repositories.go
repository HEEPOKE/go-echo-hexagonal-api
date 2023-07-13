package interfaces

import "github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"

type AuthRepository interface {
	Register(user *models.User) error
}

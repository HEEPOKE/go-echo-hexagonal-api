package services

import (
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/interfaces"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
)

type AuthService struct {
	authRepository interfaces.AuthRepository
}

func NewAuthService(authRepository interfaces.AuthRepository) *AuthService {
	return &AuthService{authRepository: authRepository}
}

func (a *AuthService) Register(user *models.User) error {
	return a.authRepository.Register(user)
}

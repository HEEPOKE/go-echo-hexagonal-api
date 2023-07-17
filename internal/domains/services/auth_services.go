package services

import (
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/interfaces"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	authRepository interfaces.AuthRepository
}

func NewAuthService(authRepository interfaces.AuthRepository) *AuthService {
	return &AuthService{authRepository: authRepository}
}

func (a *AuthService) Login(email, password string) (*models.User, error) {
	return a.authRepository.Login(email, password)
}

func (a *AuthService) Register(user *models.User) error {
	return a.authRepository.Register(user)
}

func (a *AuthService) Logout(token string) error {
	return a.authRepository.Logout(token)
}

func (a *AuthService) GenerateAccessToken(user *models.User, tokenExpiry time.Duration) (string, error) {
	return a.authRepository.GenerateAccessToken(user, tokenExpiry)
}

func (a *AuthService) VerifyAccessToken(tokenString string) (*jwt.Token, error) {
	return a.authRepository.VerifyAccessToken(tokenString)
}

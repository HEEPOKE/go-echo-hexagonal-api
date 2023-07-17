package interfaces

import (
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"github.com/golang-jwt/jwt/v5"
)

type AuthRepository interface {
	Login(email, password string) (*models.User, error)
	Register(user *models.User) error
	Logout(token string) error
	GenerateToken(user *models.User, tokenExpiry time.Duration) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}

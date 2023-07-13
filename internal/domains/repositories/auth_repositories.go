package repositories

import (
	"fmt"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// func (r *AuthRepository) Login() (*models.User, error) {
// 	var users models.User
// 	password, err := utils.Decrypt(users.Password)

// 	if err != nil {
// 		return nil, err
// 	}

// }

func (r *AuthRepository) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = string(hashedPassword)
	err = r.db.Create(user).Error
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}
	return nil
}

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

func (r *AuthRepository) Login(email, password string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	return &user, nil
}

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

func (r *AuthRepository) Logout(token string) error {
	// if err := tokenManager.InvalidateToken(token); err != nil {
	//     return fmt.Errorf("failed to logout: %w", err)
	// }
	return nil
}

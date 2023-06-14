package repositories

import (
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

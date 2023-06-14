package repositories

import (
	"errors"
	"fmt"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmailOrUsername(email, username string) (*models.User, error) {
	var user models.User
	db := r.db.Select("id, email, username, tel, role, created_at, updated_at", "deleted_at").Where("email = ? OR username = ?", email, username).First(&user)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found: %w", db.Error)
		}
		return nil, fmt.Errorf("failed to get user: %w", db.Error)
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

package repositories

import (
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"gorm.io/gorm"
)

type ScheduRepository struct {
	db *gorm.DB
}

func NewScheduRepository(db *gorm.DB) *ScheduRepository {
	return &ScheduRepository{db: db}
}

func (r *ScheduRepository) List() ([]*models.Schedu, error) {
	var schedus []*models.Schedu
	err := r.db.Find(&schedus).Error
	if err != nil {
		return nil, err
	}
	return schedus, nil
}

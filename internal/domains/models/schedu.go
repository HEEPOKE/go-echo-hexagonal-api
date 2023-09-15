package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedu struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Content    string         `gorm:"type:VARCHAR(255)" json:"content"`
	ScheduTime time.Time      `gorm:"type:DATETIME" json:"schedu_time"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

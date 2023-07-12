package models

import (
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/enums"
	"gorm.io/gorm"
)

type User struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"type:VARCHAR(255);unique" json:"email"`
	Username  string         `gorm:"type:VARCHAR(255);unique" json:"username"`
	Password  string         `gorm:"type:VARCHAR(255)" json:"password"`
	Tel       string         `gorm:"type:VARCHAR(255)" json:"tel"`
	Role      enums.Role     `gorm:"type:ENUM('ADMIN', 'USER')" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

package models

import (
	"time"

	"github.com/HEEPOKE/go-gin-hexagonal-api/pkg/enums"
	"gorm.io/gorm"
)

type User struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"type:VARCHAR(255);unique" json:"email"`
	Username  string         `gorm:"type:VARCHAR(255);unique" json:"username"`
	Password  string         `gorm:"type:VARCHAR(255)" json:"password"`
	Tel       string         `gorm:"type:VARCHAR(255)" json:"tel"`
	Role      enums.Role     `gorm:"type:ENUM('admin', 'user');default:'user'" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

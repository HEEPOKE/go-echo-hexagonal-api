package models

import "github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"

type User struct {
	models.DefaultModel
	Email    string `gorm:"type:VARCHAR(255);unique" json:"email"`
	Username string `gorm:"type:VARCHAR(255);unique" json:"username"`
	Password string `gorm:"type:VARCHAR(255)" json:"password"`
	Tel      string `gorm:"type:VARCHAR(255)" json:"tel"`
	Role     string `gorm:"type:VARCHAR(255)" json:"role"`
}

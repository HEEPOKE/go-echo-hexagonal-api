package models

type Auth struct {
	Email    string `gorm:"type:VARCHAR(255);unique" json:"email"`
	Username string `gorm:"type:VARCHAR(255);unique" json:"username"`
	Password string `gorm:"type:VARCHAR(255)" json:"password"`
}

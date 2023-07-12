package models

type Auth struct {
	Email    string `gorm:"type:VARCHAR(255);unique" json:"email"`
	Username string `gorm:"type:VARCHAR(255);unique" json:"username"`
	Password string `gorm:"type:VARCHAR(255)" json:"password"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

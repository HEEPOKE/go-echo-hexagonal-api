package models

type Auth struct {
	EmailOrUsername string `json:"email_or_username"`
	Password        string `json:"password"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

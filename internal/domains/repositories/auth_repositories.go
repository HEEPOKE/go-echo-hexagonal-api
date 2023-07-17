package repositories

import (
	"fmt"
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/constants"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db              *gorm.DB
	accessTokenKey  []byte
	refreshTokenKey []byte
}

func NewAuthRepository(db *gorm.DB, accessTokenKey, refreshTokenKey string) *AuthRepository {
	return &AuthRepository{
		db:              db,
		accessTokenKey:  []byte(accessTokenKey),
		refreshTokenKey: []byte(refreshTokenKey),
	}
}

func (r *AuthRepository) Login(email, password string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	return &user, nil
}

func (r *AuthRepository) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = string(hashedPassword)
	err = r.db.Create(user).Error
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}
	return nil
}

func (r *AuthRepository) Logout(token string) error {
	return nil
}

func (r *AuthRepository) GenerateAccessToken(user *models.User, tokenExpiry time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["username"] = user.Username
	claims["tel"] = user.Tel
	claims["role"] = user.Role
	claims["name"] = constants.ACCESS_TOKEN
	claims["exp"] = time.Now().Add(tokenExpiry).Unix()

	tokenString, err := token.SignedString(r.accessTokenKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate access_token: %w", err)
	}

	return tokenString, nil
}

func (r *AuthRepository) GenerateRefreshToken(user *models.User, tokenExpiry time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["username"] = user.Username
	claims["tel"] = user.Tel
	claims["role"] = user.Role
	claims["name"] = constants.REFRESH_TOKEN
	claims["exp"] = time.Now().Add(tokenExpiry).Unix()

	tokenString, err := token.SignedString(r.accessTokenKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh_token: %w", err)
	}

	return tokenString, nil
}

func (r *AuthRepository) VerifyAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid access_token signing method")
		}
		return []byte(r.accessTokenKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse access_token: %w", err)
	}

	return token, nil
}

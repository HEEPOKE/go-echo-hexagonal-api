package repositories

import (
	"fmt"
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db        *gorm.DB
	secretKey []byte
}

func NewAuthRepository(db *gorm.DB, secretKey string) *AuthRepository {
	return &AuthRepository{
		db:        db,
		secretKey: []byte(secretKey),
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

func (r *AuthRepository) GenerateToken(user *models.User, tokenExpiry time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["username"] = user.Username
	claims["tel"] = user.Tel
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(tokenExpiry).Unix()

	tokenString, err := token.SignedString(r.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

func (r *AuthRepository) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return r.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	return token, nil
}

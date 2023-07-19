package handlers

import (
	"net/http"
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/utils"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/services"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/constants"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/enums"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService  services.AuthService
	userService  services.UserService
	jwtSecretKey string
}

type LoginInput struct {
	EmailOrUsername string `json:"email_or_username" validate:"required" example:"heepoke"`
	Password        string `json:"password" validate:"required" example:"64765555"`
}

type RegisterInput struct {
	Email    string     `json:"email" validate:"required" example:"aaa@gmail.com"`
	Username string     `json:"username" validate:"required" example:"heepoke"`
	Password string     `json:"password" validate:"required" example:"64765555"`
	Tel      string     `json:"tel" validate:"required" example:"0000000000"`
	Role     enums.Role `json:"role" validate:"required"`
}

func NewAuthHandler(authService services.AuthService, userService services.UserService, jwtSecretKey string) *AuthHandler {
	return &AuthHandler{authService: authService, userService: userService, jwtSecretKey: jwtSecretKey}
}

// Post Login godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/login [post]
// @param Authorization header string true "Bearer token"
// @param Body body handlers.LoginInput true "login"
func (ah *AuthHandler) LoginHandler(c echo.Context) error {
	var auth models.Auth

	err := c.Bind(&auth)
	if err := c.Bind(&auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to decode user data")
	}

	user, err := ah.authService.Login(auth.EmailOrUsername, auth.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Login")
	}

	tokenString, err := ah.authService.GenerateAccessToken(user, time.Hour*24)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate access_token")
	}

	refreshTokenString, err := ah.authService.GenerateRefreshToken(user, time.Hour*72)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate refresh_token")
	}

	response := map[string]interface{}{
		"user":          user,
		"access_token":  tokenString,
		"refresh_token": refreshTokenString,
	}

	return c.JSON(http.StatusOK, response)
}

// Post Register godoc
// @Summary Register
// @Description Register
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/register [post]
// @param Authorization header string true "Bearer token"
// @param Body body handlers.RegisterInput true "Register"
func (ah *AuthHandler) RegisterHandler(c echo.Context) error {
	var user models.User

	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to decode user data")
	}

	err = ah.authService.Register(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to register")
	}

	response := map[string]interface{}{
		"message": "User registered successfully",
	}

	return c.JSON(http.StatusCreated, response)
}

// Get Logout godoc
// @Summary Logout
// @Description Logout
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/logout [get]
// @param Authorization header string true "Bearer token"
func (ah *AuthHandler) LogoutHandler(c echo.Context) error {
	token := utils.ExtractTokenFromHeader(c)
	if token == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
	}

	err := ah.authService.Logout(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to logout")
	}

	response := map[string]interface{}{
		"message": "Logout successful",
	}

	return c.JSON(http.StatusOK, response)
}

// Get Refreah Token godoc
// @Summary Refreah Token
// @Description Refreah Token
// @Tags Auth
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/refresh-token [get]
// @param Authorization header string true "Bearer token"
func (ah *AuthHandler) RefreshTokenHandler(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	userID := claims["user_id"].(int)

	user, err := ah.userService.GetUserByID(userID)
	if err != nil {
		return utils.ResponseFailOnError(c, err, constants.ERR_REFRESH_TOKEN_FAILED, http.StatusInternalServerError)
	}

	accessToken, err := ah.authService.GenerateAccessToken(user, time.Hour*24)
	if err != nil {
		return utils.ResponseFailOnError(c, err, constants.ERR_ACCESS_TOKEN_FAILED, http.StatusInternalServerError)
	}

	refreshToken, err := ah.authService.GenerateRefreshToken(user, time.Hour*72)
	if err != nil {
		return utils.ResponseFailOnError(c, err, constants.ERR_REFRESH_TOKEN_FAILED, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return c.JSON(http.StatusOK, response)
}

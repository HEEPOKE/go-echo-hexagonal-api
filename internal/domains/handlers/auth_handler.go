package handlers

import (
	"net/http"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/services"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/enums"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService services.AuthService
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

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
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

	return c.JSON(http.StatusOK, user)

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

	return c.NoContent(http.StatusOK)
}

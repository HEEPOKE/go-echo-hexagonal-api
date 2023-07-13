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

type AuthInput struct {
	Email    string     `json:"email" validate:"required" example:"aaa@gmail.com"`
	Username string     `json:"username" validate:"required" example:"heepoke"`
	Password string     `json:"password" example:"64765555"`
	Tel      string     `json:"tel" validate:"required" example:"0000000000"`
	Role     enums.Role `json:"role" validate:"required"`
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
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
// @param Body body handlers.AuthInput true "Register"
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

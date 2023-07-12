package handlers

import (
	"net/http"
	"strconv"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/services"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/enums"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService services.UserService
}

type UserInput struct {
	Email    string     `json:"email" validate:"required" example:"aaa@gmail.com"`
	Username string     `json:"username" validate:"required" example:"heepoke"`
	Password string     `json:"password" example:"yoyo5555"`
	Tel      string     `json:"tel" validate:"required" example:"0000000000"`
	Role     enums.Role `json:"role" validate:"required"`
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Get All User godoc
// @Summary Get list all Users
// @Description Get list all Users
// @Tags Users
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/all [get]
// @param Authorization header string true "Bearer token"
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get users")
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	idParam := c.QueryParam("id")
	if idParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing user ID")
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserByEmailOrUsername(c echo.Context) error {
	emailOrUsername := c.QueryParam("email_or_username")
	if emailOrUsername == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing email or username")
	}

	user, err := h.userService.GetByEmailOrUsername(emailOrUsername, emailOrUsername)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user")
	}

	return c.JSON(http.StatusOK, user)
}

// Post Create User godoc
// @Summary Create Users
// @Description Create Users
// @Tags Users
// @Accept application/json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /users/create [post]
// @param Authorization header string true "Bearer token"
// @param Body body handlers.UserInput true "CreateUsers"
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user models.User

	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to decode user data")
	}

	err = h.userService.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	return c.NoContent(http.StatusCreated)
}

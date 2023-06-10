package handlers

import (
	"net/http"
	"strconv"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/services"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

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

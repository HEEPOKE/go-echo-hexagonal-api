package handlers

import (
	"net/http"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/utils"
	"github.com/labstack/echo/v4"
)

type JwtHandler struct {
}

func NewJwtHandler() *JwtHandler {
	return &JwtHandler{}
}

func (jh *JwtHandler) AuthError(c echo.Context, e error) error {
	if e.Error() == "Token is expired" {
		return utils.ResponseFailOnError(c, e, "Unauthorized", http.StatusUnprocessableEntity)
	} else {
		return utils.ResponseFailOnError(c, e, "Unauthorized", http.StatusUnauthorized)
	}
}

func (jh *JwtHandler) AuthSuccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

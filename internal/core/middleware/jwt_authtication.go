package middleware

import (
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWT([]byte(config.Cfg.JWT_ACCESS_KEY))
}

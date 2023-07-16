package utils

import "github.com/labstack/echo/v4"

func ExtractTokenFromHeader(c echo.Context) string {
	authHeader := c.Request().Header.Get("Authorization")
	if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		return ""
	}
	return authHeader[7:]
}

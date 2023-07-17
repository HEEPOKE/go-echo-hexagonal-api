package utils

import "github.com/labstack/echo/v4"

func ResponseFailOnError(c echo.Context, err error, msg string, status int) error {
	if err != nil {
		return c.JSON(status, echo.Map{
			"msg":   msg,
			"error": err.Error(),
		})
	}
	return nil
}

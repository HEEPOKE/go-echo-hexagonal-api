package main

import (
	"github.com/HEEPOKE/go-gin-hexagonal-api/pkg/database"
	"github.com/labstack/echo/v4"
)

func main() {
	database.ConnectAndCloseDatabase()

	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}

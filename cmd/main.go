package main

import (
	"fmt"
	"log"

	"github.com/HEEPOKE/go-gin-hexagonal-api/pkg/config"
	"github.com/HEEPOKE/go-gin-hexagonal-api/pkg/database"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	database.ConnectAndCloseDatabase(cfg)

	e := echo.New()
	address := fmt.Sprintf(":%s", cfg.DBPort)
	e.Logger.Fatal(e.Start(address))
}

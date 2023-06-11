package main

import (
	"fmt"
	"log"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/repositories"
	server "github.com/HEEPOKE/go-echo-hexagonal-api/internal/http"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/config"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/database"
)

// @title Swagger Example API
// @title Go Echo Hexagonal API
// @version 1.0
// @description This is a Go Echo Hexagonal API server.
// @host localhost:6476
// @BasePath /api
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := database.ConnectAndCloseDatabase(cfg)

	userRepository := repositories.NewUserRepository(db)

	address := fmt.Sprintf(":%s", cfg.PORT)
	http := server.NewServer(userRepository)
	http.RouteInit(address)
}

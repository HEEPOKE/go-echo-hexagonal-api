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
// @contact.name API Support
// @contact.url https://github.com/HEEPOKE
// @contact.email damon1FX@gmail.com
// @host localhost:6476
// @BasePath /apis
// @schemes http https
func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repositories.NewUserRepository(db)
	authRepository := repositories.NewAuthRepository(db, config.Cfg.JWT_SECRET_KEY)

	address := fmt.Sprintf(":%s", config.Cfg.PORT)
	http := server.NewServer(userRepository, authRepository, config.Cfg.JWT_SECRET_KEY)
	http.RouteInit(address)
}

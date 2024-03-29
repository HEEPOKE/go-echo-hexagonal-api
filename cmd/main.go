package main

import (
	"fmt"
	"log"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/app/tasks"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/repositories"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/services"
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
// @contact.email Damon1FX@gmail.com
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

	database.CheckRedis()

	authRepository := repositories.NewAuthRepository(db, config.Cfg.JWT_ACCESS_KEY, config.Cfg.JWT_REFRESH_KEY)
	userRepository := repositories.NewUserRepository(db)
	scheduRepository := repositories.NewScheduRepository(db)

	scheduService := services.NewScheduService(scheduRepository)
	scheduServiceValue := *scheduService

	taskAuto := tasks.NewScheduTask(scheduServiceValue)
	taskAuto.AutoTaskRunner()

	address := fmt.Sprintf(":%s", config.Cfg.PORT)
	http := server.NewServer(userRepository, authRepository, scheduRepository, config.Cfg.JWT_ACCESS_KEY, config.Cfg.JWT_REFRESH_KEY)
	http.RouteInit(address)
}

package main

import (
	"fmt"
	"log"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/repositories"
	server "github.com/HEEPOKE/go-echo-hexagonal-api/internal/http"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/config"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/database"
)

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

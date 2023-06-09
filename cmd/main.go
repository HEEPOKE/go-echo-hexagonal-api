package main

import (
	"fmt"
	"log"

	server "github.com/HEEPOKE/go-gin-hexagonal-api/internal/http"
	"github.com/HEEPOKE/go-gin-hexagonal-api/pkg/config"
	"github.com/HEEPOKE/go-gin-hexagonal-api/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	database.ConnectAndCloseDatabase(cfg)

	address := fmt.Sprintf(":%s", cfg.PORT)
	http := server.NewServer()
	http.RouteInit(address)
}

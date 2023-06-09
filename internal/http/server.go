package http

import (
	"log"

	"github.com/HEEPOKE/go-gin-hexagonal-api/internal/core/interfaces"
	"github.com/HEEPOKE/go-gin-hexagonal-api/internal/domains/handlers"
	"github.com/HEEPOKE/go-gin-hexagonal-api/internal/domains/services"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo        *echo.Echo
	userHandler *handlers.UserHandler
}

func NewServer(userRepository interfaces.UserRepository) *Server {
	e := echo.New()

	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(*userService)

	return &Server{
		echo:        e,
		userHandler: userHandler,
	}
}

func (s *Server) RouteInit(address string) {
	s.routeConfig()

	err := s.echo.Start(address)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func (s *Server) routeConfig() {
	api := s.echo.Group("/api")

	api.GET("/users/all", s.userHandler.GetAllUsers)
	api.GET("/users/find/:id", s.userHandler.GetUserByID)
	api.POST("/users/create", s.userHandler.CreateUser)
}

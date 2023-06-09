package http

import (
	"log"
	"net/http"

	"github.com/HEEPOKE/go-gin-hexagonal-api/internal/core/interfaces"
	"github.com/HEEPOKE/go-gin-hexagonal-api/internal/domains/handlers"
	"github.com/HEEPOKE/go-gin-hexagonal-api/internal/domains/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo        *echo.Echo
	userHandler *handlers.UserHandler
}

func NewServer(userRepository interfaces.UserRepository) *Server {
	e := echo.New()

	loggerConfig := middleware.LoggerConfig{
		Format:           "URI=${uri}\n, METHOD=${method},  STATUS=${status}, HEADER=${header}\n, QUERY=${query}\n, ERROR=${error}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.LoggerWithConfig(loggerConfig))
	e.Use(middleware.Recover())

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

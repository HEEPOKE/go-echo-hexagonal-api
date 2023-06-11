package http

import (
	"log"
	"net/http"

	_ "github.com/HEEPOKE/go-echo-hexagonal-api/internal/app/docs"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/interfaces"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/handlers"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	echo        *echo.Echo
	userHandler *handlers.UserHandler
}

func NewServer(userRepository interfaces.UserRepository) *Server {
	e := echo.New()

	loggerConfig := middleware.LoggerConfig{
		Format:           "URI::${uri}\n, METHOD::${method},  STATUS::${status}, HEADER::${header}\n, QUERY::${query}\n, ERROR::${error}\n",
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

// @title Swagger Example API
// @version 1.0
// @description This is a API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:6476
// @BasePath /api
func (s *Server) RouteInit(address string) {
	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
	s.routeConfig()

	err := s.echo.Start(address)
	if err != nil {
		log.Fatalf("Failed To Start The Server: %v", err)
	}
}

func (s *Server) routeConfig() {

	api := s.echo.Group("/api")

	api.GET("/users/all", s.userHandler.GetAllUsers)
	api.GET("/users/find/:id", s.userHandler.GetUserByID)
	api.POST("/users/create", s.userHandler.CreateUser)
}

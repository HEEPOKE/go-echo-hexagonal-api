package http

import (
	"log"
	"net/http"

	_ "github.com/HEEPOKE/go-echo-hexagonal-api/internal/app/docs"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/interfaces"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/handlers"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/services"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	echo        *echo.Echo
	userHandler *handlers.UserHandler
	authHandler *handlers.AuthHandler
}

func NewServer(userRepository interfaces.UserRepository, authRepository interfaces.AuthRepository, jwtSecretKey string) *Server {
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

	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(*authService, jwtSecretKey)

	return &Server{
		echo:        e,
		userHandler: userHandler,
		authHandler: authHandler,
	}
}

func (s *Server) RouteInit(address string) {
	s.routeConfig()

	err := s.echo.Start(address)
	if err != nil {
		log.Fatalf("Failed To Start The Server: %v", err)
	}
}

func (s *Server) routeConfig() {
	api := s.echo.Group("/apis")

	api.GET("/docs/*", echoSwagger.WrapHandler)

	jwtMiddleware := echoJwt.WithConfig(echoJwt.Config{
		SigningKey: []byte("jwt-secret-key"),
	})

	user := api.Group("/users")

	user.Use(jwtMiddleware)

	user.GET("/all", s.userHandler.GetAllUsers)
	user.GET("/find/:id", s.userHandler.GetUserByID)
	user.GET("/find/:email_or_username", s.userHandler.GetUserByEmailOrUsername)
	user.POST("/create", s.userHandler.CreateUser)

	auth := api.Group("/auth")

	auth.POST("/login", s.authHandler.LoginHandler)
	auth.POST("/register", s.authHandler.RegisterHandler)
	auth.GET("/logout", s.authHandler.LogoutHandler)
}

package http

import (
	"log"
	"net/http"

	_ "github.com/HEEPOKE/go-echo-hexagonal-api/internal/app/docs"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/interfaces"
	myMidleware "github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/middleware"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/utils"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/handlers"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/services"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	echo        *echo.Echo
	jwtHandler  *handlers.JwtHandler
	authHandler *handlers.AuthHandler
	userHandler *handlers.UserHandler
}

func NewServer(userRepository interfaces.UserRepository, authRepository interfaces.AuthRepository, jwtSecretKey string) *Server {
	e := echo.New()

	loggerConfig := middleware.LoggerConfig{
		Format:           "URI::${uri}\n, METHOD::${method},  STATUS::${status}, HEADER::${header}\n, QUERY::${query}\n, ERROR::${error}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           utils.ColorLoggerOutput(),
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.LoggerWithConfig(loggerConfig))
	e.Use(middleware.Recover())
	e.Use(myMidleware.JWTMiddleware())

	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(*userService)

	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(*authService, jwtSecretKey)

	jwtHandler := handlers.NewJwtHandler()

	return &Server{
		echo:        e,
		userHandler: userHandler,
		authHandler: authHandler,
		jwtHandler:  jwtHandler,
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

	user := api.Group("/users", echojwt.JWT([]byte(config.Cfg.JWT_ACCESS_KEY)))

	user.GET("/all", s.userHandler.GetAllUsers)
	user.GET("/find/:id", s.userHandler.GetUserByID)
	user.GET("/find/:email_or_username", s.userHandler.GetUserByEmailOrUsername)
	user.POST("/create", s.userHandler.CreateUser)

	auth := api.Group("/auth")

	auth.POST("/login", s.authHandler.LoginHandler)
	auth.POST("/register", s.authHandler.RegisterHandler)
	auth.GET("/logout", s.authHandler.LogoutHandler)
}

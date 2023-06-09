package http

import (
	"log"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	return &Server{
		echo: echo.New(),
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
}

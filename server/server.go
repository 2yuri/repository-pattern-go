package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/repository-pattern-go/config"
	"github.com/hyperyuri/repository-pattern-go/server/routes"
)

type Server struct {
	host   string
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		host:   config.GetConfig().ServerHost,
		port:   config.GetConfig().ServerPort,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}

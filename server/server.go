package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/repository-pattern-go/routes"
)

type Server struct {
	host   string
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		host:   os.Getenv("HOST"),
		port:   os.Getenv("PORT"),
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Fatal(router.Run(":" + s.port))
}

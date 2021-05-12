package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/repository-pattern-go/database"
	"github.com/hyperyuri/repository-pattern-go/server"
)

type Server struct {
	host   string
	port   string
	server *gin.Engine
}

func main() {
	database.StartDatabase()

	s := server.NewServer()

	s.Run()
}

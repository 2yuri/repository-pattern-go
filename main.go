package main

import (
	"github.com/hyperyuri/repository-pattern-go/database"
	"github.com/hyperyuri/repository-pattern-go/server"
)

func main() {
	database.StartDatabase()

	s := server.NewServer()

	s.Run()
}

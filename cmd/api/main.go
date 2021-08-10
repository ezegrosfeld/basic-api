package main

import (
	"github.com/ezegrosfeld/basic-api/cmd/api/server"
	"github.com/ezegrosfeld/basic-api/pkg/db"
	"github.com/ezegrosfeld/basic-api/pkg/logger"
)

func main() {
	db.InitializeDatabase()
	defer db.Database.Close()

	logger.Initialize()
	server.StartServer()
}

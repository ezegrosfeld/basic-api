package server

import (
	"log"

	"github.com/ezegrosfeld/basic-api/internal/health"
	"github.com/ezegrosfeld/basic-api/pkg/http"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func initServer() {
	// Create gin router
	router = http.Router

	// Register repository
	repo := health.NewRepository()
	// Register service
	service := health.NewService(repo)
	// Register controller
	controller := health.NewController(service)

	// Register routes
	router.GET("/health", controller.Get())
}

func StartServer() {
	// initialize server
	initServer()

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Panic(err)
	}
}

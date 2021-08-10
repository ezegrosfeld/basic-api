package http

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	Router     *gin.Engine
	production bool = os.Getenv("ENVIRONMENT") == "production"
)

func init() {
	if production {
		gin.SetMode(gin.ReleaseMode)
	}

	Router = gin.New()

	if !production {
		Router.Use(gin.Logger())
	}

	Router.Use(gin.Recovery())
}

package health

import (
	"github.com/ezegrosfeld/basic-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Controller struct {
	service Service
	log     zerolog.Logger
}

func NewController(service Service) *Controller {
	log := logger.New("health", "controller")
	return &Controller{
		service: service,
		log:     log,
	}
}

func (c *Controller) Get() gin.HandlerFunc {
	c.log.Debug().Msg("Answering health check")
	return func(ctx *gin.Context) {
		h, err := c.service.Get()
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, h)
	}
}

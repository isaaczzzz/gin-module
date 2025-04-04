package api

import (
	"github.com/gin-gonic/gin"
	"github.com/isaaczzzz/gin-module/api/echo"
	"github.com/isaaczzzz/gin-module/internal/service"
)

type Controllers struct {
	EchoController *echo.EchoController
}

func InitControllers(services *service.Services) *Controllers {
	return &Controllers{
		EchoController: echo.NewEchoController(services.EchoService),
	}
}

func RegisterRouters(r *gin.Engine, controllers *Controllers) {
	echo.RegisterRoutes(r, controllers.EchoController)
}

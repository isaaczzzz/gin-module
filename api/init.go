package api

import (
	"github.com/gin-gonic/gin"
	"github.com/isaaczzzz/gin-module/api/echo"
	"github.com/isaaczzzz/gin-module/pkg/rpcclient"
)

type Controllers struct {
	EchoController *echo.EchoController
}

func InitControllers(rpcClients *rpcclient.RPCClients) *Controllers {
	return &Controllers{
		EchoController: echo.NewEchoController(rpcClients.EchoClient),
	}
}

func RegisterRouters(r *gin.Engine, controllers *Controllers) {
	echo.RegisterRoutes(r, controllers.EchoController)
}

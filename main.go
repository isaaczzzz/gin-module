package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/isaaczzzz/gin-module/api"
	"github.com/isaaczzzz/gin-module/pkg/config"
	"github.com/isaaczzzz/gin-module/pkg/rpcclient"
	"log"
	"net/http"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化 gRPC 连接
	rpcClients := rpcclient.InitRPCClients(cfg)
	defer func(rpcClients *rpcclient.RPCClients) {
		err := rpcClients.CloseRPCClients()
		if err != nil {
			log.Fatalf("Failed to close rpc clients: %v", err)
		}
	}(rpcClients)

	// 创建 Controller 实例
	controllers := api.InitControllers(rpcClients)

	// 创建 Gin 引擎
	r := gin.Default()

	// 注册 echo 路由
	api.RegisterRouters(r, controllers)

	// 启动 HTTP 服务
	err = r.Run(cfg.HTTP.Port)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Failed to start server: %v", err)
	}
}

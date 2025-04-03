package echo

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/isaaczzzz/gin-module/pkg/rpcclient"
	"github.com/isaaczzzz/gin-module/rpc_gen/api/echo"
	"log"
	"net/http"
	"time"
)

// EchoController 处理 Echo 请求
type EchoController struct {
	echoClient echo.EchoServiceClient
}

// NewEchoController 创建一个新的 EchoController 实例
func NewEchoController(client *rpcclient.EchoClient) *EchoController {
	return &EchoController{
		echoClient: echo.NewEchoServiceClient(client.Conn),
	}
}

// Echo 处理 /api/echo/echo 请求
func (c *EchoController) Echo(ctx *gin.Context) {
	message := ctx.Query("message")
	if message == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "message parameter is required"})
		return
	}

	grpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &echo.EchoRequest{
		Message: message,
	}

	resp, err := c.echoClient.Echo(grpcCtx, req)
	if err != nil {
		log.Printf("Failed to call gRPC Echo service: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call gRPC Echo service"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": resp.Message})
}

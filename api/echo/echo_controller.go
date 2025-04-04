package echo

import (
	"github.com/gin-gonic/gin"
	"github.com/isaaczzzz/gin-module/internal/service"
	"net/http"
)

// EchoController 处理 Echo 请求
type EchoController struct {
	ehcoService service.IEchoService
}

// NewEchoController 创建一个新的 EchoController 实例
func NewEchoController(service service.IEchoService) *EchoController {
	return &EchoController{
		ehcoService: service,
	}
}

// Echo 处理 /api/echo/echo 请求
func (c *EchoController) Echo(ctx *gin.Context) {
	message := ctx.Query("message")
	if message == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "message parameter is required"})
		return
	}

	message, err := c.ehcoService.Echo(ctx, message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

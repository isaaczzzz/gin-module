package echo

import "github.com/gin-gonic/gin"

// RegisterRoutes 注册 echo 相关的路由
func RegisterRoutes(r *gin.Engine, controller *EchoController) {
	accessGroup := r.Group("/api")
	{
		accessGroup.GET("/echo", controller.Echo)
	}
}
